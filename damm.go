package damm

type Damm interface {
	Generate(digits []int) int
	Verify(digits []int) bool
	Modulus() int
}

func calculate(digits []int, modulus, mask int) (checkDigit int) {
	for _, digit := range digits {
		checkDigit ^= digit
		checkDigit <<= 1
		if checkDigit >= modulus {
			checkDigit ^= mask
		}
	}
	return checkDigit
}

type damm struct {
	modulus int
	mask    int
}

func (d *damm) Generate(digits []int) int {
	return calculate(digits, d.modulus, d.mask)
}

func (d *damm) Verify(digits []int) bool {
	return calculate(digits, d.modulus, d.mask) == 0
}

func (d *damm) Modulus() int {
	return d.modulus
}

func New32() Damm {
	const modulus = 32
	return &damm{
		modulus: modulus,
		mask:    modulus | 5,
	}
}

func New64() Damm {
	const modulus = 64
	return &damm{
		modulus: modulus,
		mask:    modulus | 3,
	}
}
