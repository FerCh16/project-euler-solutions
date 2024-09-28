package projecteulersolutions

/*
Problem 3
~~~~~~~~~~
The prime factors of 13195 are: 5,7,13,29.

What is the largest prime factor of the number 600851475143?
*/
func LargestPrimeFactor(n int) int {
	var greatest int

	for n%2 == 0 {
		n /= 2
		greatest = 2
	}

	for i := 3; i*i < n; i = i + 2 {
		for n%i == 0 {
			greatest = i
			n /= i
		}
	}

	if n > 2 {
		greatest = n
	}
	return greatest
}
