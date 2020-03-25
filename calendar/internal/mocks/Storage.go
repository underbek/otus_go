// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	context "context"

	domain "github.com/AndreyAndreevich/otus_go/calendar/internal/domain"
	mock "github.com/stretchr/testify/mock"

	time "time"

	uuid "github.com/google/uuid"
)

// Storage is an autogenerated mock type for the Storage type
type Storage struct {
	mock.Mock
}

// GetEventsInTime provides a mock function with given fields: ctx, _a1, duration
func (_m *Storage) GetEventsInTime(ctx context.Context, _a1 time.Time, duration time.Duration) ([]domain.Event, error) {
	ret := _m.Called(ctx, _a1, duration)

	var r0 []domain.Event
	if rf, ok := ret.Get(0).(func(context.Context, time.Time, time.Duration) []domain.Event); ok {
		r0 = rf(ctx, _a1, duration)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Event)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, time.Time, time.Duration) error); ok {
		r1 = rf(ctx, _a1, duration)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Insert provides a mock function with given fields: ctx, event
func (_m *Storage) Insert(ctx context.Context, event domain.Event) error {
	ret := _m.Called(ctx, event)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.Event) error); ok {
		r0 = rf(ctx, event)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Listing provides a mock function with given fields: ctx
func (_m *Storage) Listing(ctx context.Context) ([]domain.Event, error) {
	ret := _m.Called(ctx)

	var r0 []domain.Event
	if rf, ok := ret.Get(0).(func(context.Context) []domain.Event); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Event)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Remove provides a mock function with given fields: ctx, id
func (_m *Storage) Remove(ctx context.Context, id uuid.UUID) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Update provides a mock function with given fields: ctx, event
func (_m *Storage) Update(ctx context.Context, event domain.Event) error {
	ret := _m.Called(ctx, event)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.Event) error); ok {
		r0 = rf(ctx, event)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
