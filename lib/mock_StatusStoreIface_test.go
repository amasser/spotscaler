package autoscaler

import mock "github.com/stretchr/testify/mock"
import time "time"

// MockStatusStoreIface is an autogenerated mock type for the StatusStoreIface type
type MockStatusStoreIface struct {
	mock.Mock
}

// AddSchedules provides a mock function with given fields: sch
func (_m *MockStatusStoreIface) AddSchedules(sch *Schedule) error {
	ret := _m.Called(sch)

	var r0 error
	if rf, ok := ret.Get(0).(func(*Schedule) error); ok {
		r0 = rf(sch)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteTimer provides a mock function with given fields: key
func (_m *MockStatusStoreIface) DeleteTimer(key string) error {
	ret := _m.Called(key)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(key)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FetchCooldownEndsAt provides a mock function with given fields:
func (_m *MockStatusStoreIface) FetchCooldownEndsAt() (time.Time, error) {
	ret := _m.Called()

	var r0 time.Time
	if rf, ok := ret.Get(0).(func() time.Time); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(time.Time)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetExpiredTimers provides a mock function with given fields:
func (_m *MockStatusStoreIface) GetExpiredTimers() ([]string, error) {
	ret := _m.Called()

	var r0 []string
	if rf, ok := ret.Get(0).(func() []string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListSchedules provides a mock function with given fields:
func (_m *MockStatusStoreIface) ListSchedules() ([]*Schedule, error) {
	ret := _m.Called()

	var r0 []*Schedule
	if rf, ok := ret.Get(0).(func() []*Schedule); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*Schedule)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RemoveSchedules provides a mock function with given fields: key
func (_m *MockStatusStoreIface) RemoveSchedules(key string) error {
	ret := _m.Called(key)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(key)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// StoreCooldownEndsAt provides a mock function with given fields: t
func (_m *MockStatusStoreIface) StoreCooldownEndsAt(t time.Time) error {
	ret := _m.Called(t)

	var r0 error
	if rf, ok := ret.Get(0).(func(time.Time) error); ok {
		r0 = rf(t)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateTimer provides a mock function with given fields: key, t
func (_m *MockStatusStoreIface) UpdateTimer(key string, t time.Time) error {
	ret := _m.Called(key, t)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, time.Time) error); ok {
		r0 = rf(key, t)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

var _ StatusStoreIface = (*MockStatusStoreIface)(nil)