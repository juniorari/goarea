package goarea

import "math"

// Pi função pública
const Pi = 3.1416

// Circ função pública, calcula a circunferencia
func Circ(raio float64) float64 {
	return math.Pow(raio, 2) * Pi
}

// Rect Calcula a área
func Rect(base, altura float64) float64 {
	return base * altura
}

// Não é visível
func _TrianguloEq(base, altura float64) float64 {
	return (base * altura) / 2
}
