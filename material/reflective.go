package material

import (
	"github.com/nico-ulbricht/go-pathtracer/geometry"
)

type ReflectiveMaterial struct {
	reflectance float64
	roughness   float64
}

func NewReflectiveMaterial(reflectance, roughness float64) *ReflectiveMaterial {
	return &ReflectiveMaterial{reflectance, roughness}
}

func (mat *ReflectiveMaterial) Reflect(ray *geometry.Ray, intersection *geometry.Intersection) *geometry.Ray {
	directionDotNormal := ray.Direction.Dot(intersection.Normal)
	ray.Direction = ray.Direction.Subtract(intersection.Normal.MultiplyScalar(2. * directionDotNormal))
	ray.Origin = intersection.Point
	ray.Intensity *= mat.reflectance * .96
	return ray
}
