package geos

/*
#cgo LDFLAGS: -lgeos_c
#include "geos_c.h"
#include <stdlib.h>
*/
import "C"

type PreparedGeom struct {
	v *C.GEOSPreparedGeometry
}

func (this *Geos) Prepare(geom *Geom) *PreparedGeom {
	prep := C.GEOSPrepare_r(this.v, geom.v)
	if prep == nil {
		return nil
	}
	return &PreparedGeom{prep}
}

func (this *Geos) PreparedContains(a *PreparedGeom, b *Geom) bool {
	result := C.GEOSPreparedContains_r(this.v, a.v, b.v)
	if result == 1 {
		return true
	}
	// result == 2 -> exception (already logged to console)
	return false
}

func (this *Geos) PreparedIntersects(a *PreparedGeom, b *Geom) bool {
	result := C.GEOSPreparedIntersects_r(this.v, a.v, b.v)
	if result == 1 {
		return true
	}
	// result == 2 -> exception (already logged to console)
	return false
}
