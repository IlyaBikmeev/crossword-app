package core

type Metric interface {
	Evaluate(g *Grid) float64
}

type DensityAndIntersectionMetric struct {
	densityWeight      float64
	intesectionsWeight float64
}

type DensityMetric struct {
}

func NewDensityAndIntersectionMetric(dw, iw float64) *DensityAndIntersectionMetric {
	return &DensityAndIntersectionMetric{
		densityWeight:      dw,
		intesectionsWeight: iw,
	}
}

type ScoredGrid struct {
	Grid   *Grid
	Score  float64
	Metric Metric
}

func (m *DensityAndIntersectionMetric) Evaluate(g *Grid) float64 {
	d := g.Density()
	i := float64(g.Intersections()) / float64(g.Area())

	raw := m.densityWeight*d + m.intesectionsWeight*i

	normalized := raw / (m.densityWeight + m.intesectionsWeight)

	return normalized * 100.0
}

func (d DensityMetric) Evaluate(g *Grid) float64 {
	return g.Density()
}
