package main

// IRQs
const (
	cassetteRiseIrqMask = 1 << iota
	cassetteFallIrqMask
	timerIrqMask
	ioBusIrqMask
	uartSendIrqMask
	uartReceiveIrqMask
	uartErrorIrqMask
)

// NMIs
const (
	resetNmiMask    = 0x20 << iota
	diskMotorOffNmiMask
	diskIntrqNmiMask
)

func (cpu *cpu) setIrqMask(irqMask byte) {
	cpu.irqMask = irqMask
}

func (cpu *cpu) setNmiMask(nmiMask byte) {
	// Reset is always allowed:
	cpu.nmiMask = nmiMask | resetNmiMask
	cpu.updateNmiSeen()
}

func (cpu *cpu) updateNmiSeen() {
	if (cpu.nmiLatch&cpu.nmiMask) == 0 {
		cpu.nmiSeen = false
	}
}

func (cpu *cpu) handleIrq() {
	cpu.pushWord(cpu.pc)
	cpu.iff1 = false
	cpu.pc = 0x38
}

func (cpu *cpu) handleNmi() {
	cpu.pushWord(cpu.pc)
	cpu.iff1 = false
	cpu.pc = 0x66
}

func (cpu *cpu) resetButtonInterrupt(state bool) {
    if state {
		cpu.nmiLatch |= resetNmiMask
	} else {
		cpu.nmiLatch &^= resetNmiMask
	}
	cpu.updateNmiSeen()
}

func (cpu *cpu) diskMotorOffInterrupt(state bool) {
    if state {
		cpu.nmiLatch |= diskMotorOffNmiMask
	} else {
		cpu.nmiLatch &^= diskMotorOffNmiMask
	}
	cpu.updateNmiSeen()
}

func (cpu *cpu) diskIntrqInterrupt(state bool) {
    if state {
		cpu.nmiLatch |= diskIntrqNmiMask
	} else {
		cpu.nmiLatch &^= diskIntrqNmiMask
	}
	cpu.updateNmiSeen()
}

func (cpu *cpu) diskDrqInterrupt(state bool) {
	// No effect.
}
