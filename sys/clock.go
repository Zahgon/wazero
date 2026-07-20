package sys

type ClockResolution uint32

type Walltime func() (sec int64, nsec int32)

type Nanotime func() int64

type Nanosleep func(ns int64)

type Osyield func()
