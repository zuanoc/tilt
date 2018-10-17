package hud

import (
	"context"

	"github.com/windmilleng/tilt/internal/logger"
	"github.com/windmilleng/tilt/internal/store"

	"github.com/windmilleng/tilt/internal/hud/view"
)

var _ HeadsUpDisplay = (*FakeHud)(nil)

type FakeHud struct {
	LastView view.View
	Updates  chan view.View
}

func NewFakeHud() *FakeHud {
	return &FakeHud{
		Updates: make(chan view.View, 10),
	}
}

func (h *FakeHud) Run(ctx context.Context, st *store.Store) error { return nil }

func (h *FakeHud) SetNarrationMessage(ctx context.Context, msg string) {}

func (h *FakeHud) OnChange(ctx context.Context, st *store.Store) {
	state := st.RLockState()
	view := store.StateToView(state)
	st.RUnlockState()

	err := h.Update(view)
	if err != nil {
		logger.Get(ctx).Infof("Error updating HUD: %v", err)
	}
}

func (h *FakeHud) Update(v view.View) error {
	h.LastView = v
	h.Updates <- v
	return nil
}
