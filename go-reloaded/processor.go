package main

func Processor(w string) string {
	w = Convert(w)
		w = CommandNWord(w)
			w = FixAToAn(w)


	w = EdgeCases(w)
	return w
}
