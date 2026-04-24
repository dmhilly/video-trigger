package videotrigger

import (
	"context"
	"errors"
	"fmt"
	"time"

	"go.viam.com/rdk/logging"
	"go.viam.com/rdk/resource"
	generic "go.viam.com/rdk/services/generic"
	"go.viam.com/rdk/services/video"
	"go.viam.com/rdk/services/vision"
)

const prerollPadding = 5 * time.Second

var (
	GenericService   = resource.NewModel("devin-hilly", "video-trigger", "generic_service")
	errUnimplemented = errors.New("unimplemented")
)

func init() {
	resource.RegisterService(generic.API, GenericService,
		resource.Registration[resource.Resource, *Config]{
			Constructor: newVideoTriggerGenericService,
		},
	)
}

type Config struct {
	Camera        string  `json:"camera"`
	VisionService string  `json:"vision_service"`
	VideoService  string  `json:"video_service"`
	TriggerLabel  string  `json:"trigger_label"`
	Threshold     float64 `json:"threshold"`
}

// Validate ensures all parts of the config are valid and important fields exist.
// Returns three values:
//  1. Required dependencies: other resources that must exist for this resource to work.
//  2. Optional dependencies: other resources that may exist but are not required.
//  3. An error if any Config fields are missing or invalid.
//
// The `path` parameter indicates
// where this resource appears in the machine's JSON configuration
// (for example, "components.0"). You can use it in error messages
// to indicate which resource has a problem.
func (cfg *Config) Validate(path string) ([]string, []string, error) {
	if cfg.Camera == "" {
		return nil, nil, fmt.Errorf("%s: camera is required", path)
	}
	if cfg.VisionService == "" {
		return nil, nil, fmt.Errorf("%s: vision_service is required", path)
	}
	if cfg.VideoService == "" {
		return nil, nil, fmt.Errorf("%s: video_service is required", path)
	}
	if cfg.TriggerLabel == "" {
		return nil, nil, fmt.Errorf("%s: trigger_label is required", path)
	}
	return []string{cfg.VisionService, cfg.VideoService}, nil, nil
}

type videoTriggerGenericService struct {
	resource.AlwaysRebuild

	name resource.Name

	logger    logging.Logger
	cfg       *Config
	visionSvc vision.Service
	videoSvc  video.Service

	cancelCtx  context.Context
	cancelFunc func()
}

func newVideoTriggerGenericService(ctx context.Context, deps resource.Dependencies, rawConf resource.Config, logger logging.Logger) (resource.Resource, error) {
	conf, err := resource.NativeConfig[*Config](rawConf)
	if err != nil {
		return nil, err
	}

	return NewGenericService(ctx, deps, rawConf.ResourceName(), conf, logger)
}

func NewGenericService(ctx context.Context, deps resource.Dependencies, name resource.Name, conf *Config, logger logging.Logger) (resource.Resource, error) {
	visionSvc, err := vision.FromDependencies(deps, conf.VisionService)
	if err != nil {
		return nil, fmt.Errorf("failed to get vision service %q: %w", conf.VisionService, err)
	}

	videoSvc, err := video.FromDependencies(deps, conf.VideoService)
	if err != nil {
		return nil, fmt.Errorf("failed to get video service %q: %w", conf.VideoService, err)
	}

	cancelCtx, cancelFunc := context.WithCancel(context.Background())

	s := &videoTriggerGenericService{
		name:       name,
		logger:     logger,
		cfg:        conf,
		visionSvc:  visionSvc,
		videoSvc:   videoSvc,
		cancelCtx:  cancelCtx,
		cancelFunc: cancelFunc,
	}

	go s.monitor()

	return s, nil
}

func (s *videoTriggerGenericService) monitor() {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	active := false
	var startTime time.Time

	for {
		select {
		case <-s.cancelCtx.Done():
			return
		case <-ticker.C:
			detections, err := s.visionSvc.DetectionsFromCamera(s.cancelCtx, s.cfg.Camera, nil)
			if err != nil {
				s.logger.Warnf("failed to get detections from vision service: %v", err)
				continue
			}

			detected := false
			for _, d := range detections {
				if d.Label() == s.cfg.TriggerLabel && d.Score() >= s.cfg.Threshold {
					detected = true
					break
				}
			}

			if detected && !active {
				active = true
				startTime = time.Now().Add(-prerollPadding)
			} else if !detected && active {
				active = false
				endTime := time.Now()
				start := startTime.UTC().Format("2006-01-02_15-04-05Z")
				end := endTime.UTC().Format("2006-01-02_15-04-05Z")
				go s.processEvent(start, end, endTime)
			}
		}
	}
}

func (s *videoTriggerGenericService) processEvent(start, end string, endTime time.Time) {
	if wait := time.Until(endTime) + 45*time.Second; wait > 0 {
		time.Sleep(wait)
	}
	if _, err := s.videoSvc.DoCommand(s.cancelCtx, map[string]interface{}{"command": "save", "from": start, "to": end}); err != nil {
		s.logger.Warnf("failed to send DoCommand to video service: %v", err)
	}
}

func (s *videoTriggerGenericService) Name() resource.Name {
	return s.name
}

func (s *videoTriggerGenericService) DoCommand(ctx context.Context, cmd map[string]interface{}) (map[string]interface{}, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *videoTriggerGenericService) Close(context.Context) error {
	s.cancelFunc()
	return nil
}
