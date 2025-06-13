package instrucciones

type InstruccionesContexto struct {
	// The console is the output of the REPL
	Console *Console
	// The scope is the current scope of the REPL
	RegistroAmbito *RegistroAmbito
	// The call stack is the stack of breakable, continueable and returnable items
	PilaLlamada *PilaLlamada
	// Error table is the table of errors
	TablaError *TablaError
}
