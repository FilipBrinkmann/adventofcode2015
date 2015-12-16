"""Solver for advent of code no.7"""


class Gate(object):
    """Represents a gate object."""

    ASSIGN = '->'

    def __init__(self, description):
        super(Gate, self).__init__()
        self.wire = self._extractWire(description)
        self.gate_type = self._extract_gate_type(description)
        self.inputs = self._extract_inputs(description)
        print self.inputs
        # print self.inputs
        self.output = None
        self.ready_to_compute = self.can_compute()
        WIRES[self.wire] = self.output
        for inp in self.inputs:
            if inp not in WIRES.keys():
                WIRES[inp] = None

    def _extractWire(self, description):
        return description.split()[len(description.split())-1]

    def _extract_gate_type(self, description):
        parts = description.split()
        if parts[0] == 'NOT':
            return parts[0]
        elif parts[1] == self.ASSIGN:
            return 'ASSIGN'
        else:
            if parts[1] in ['AND', 'OR', 'LSHIFT', 'RSHIFT']:
                return parts[1]
            else:
                raise ValueError("Unknown command")

    def _extract_inputs(self, description):
        assert self.gate_type is not None
        # print self.description.split()
        if self.gate_type == 'ASSIGN':
            return [description.split()[0]]
        elif self.gate_type == 'NOT':
            return [description.split()[1]]
        elif self.gate_type in ['AND', 'OR', 'LSHIFT', 'RSHIFT']:
            return [description.split()[0], description.split()[2]]
        else:
            raise Exception("Gate type is None!")

    def _is_number(self, thing):
        if thing is None:
            return False
        try:
            int(thing)
            return True
        except (TypeError, ValueError):
            return False

    def _update_inputs(self):
        # print self.gate_type + " " + self.inputs.__str__()
        for i in range(len(self.inputs)):
            if not self._is_number(self.inputs[i]) and self.inputs[i] in WIRES.keys():
                self.inputs[i] = WIRES[self.inputs[i]]
        for i in range(len(self.inputs)):
            if not self._is_number(self.inputs[i]):
                self.ready_to_compute = False
                return
            else:
                self.inputs[i] = int(self.inputs[i])
                self.ready_to_compute = True
        if self.ready_to_compute is True:
            print "Gate ready to compute."

    def can_compute(self):
        self._update_inputs()
        return self.ready_to_compute

    def execute(self):
        if self.output is not None:
            raise Exception("you must not call execute() more than once")

        if not self.can_compute():
            raise Exception("error: gate is not ready to compute yet. First, provide all ou")
        if self.gate_type == 'ASSIGN':
            self.output = self.inputs[0]
        elif self.gate_type == 'NOT':
            self.output = ~ self.inputs[0]
        elif self.gate_type == 'AND':
            self.output = self.inputs[0] & self.inputs[1]
        elif self.gate_type == 'OR':
            self.output = self.inputs[0] | self.inputs[1]
        elif self.gate_type == 'LSHIFT':
            self.output = self.inputs[0] << self.inputs[1]
        elif self.gate_type == 'RSHIFT':
            self.output = self.inputs[0] >> self.inputs[1]
        else:
            raise Exception('unknown gate_type')

    def get_wirename(self):
        return self.wire

    def get_output(self):
        return self.output

WIRES = {}
UNFINISHED_GATES = []

def value_of(wire):
    return WIRES[wire]

def process_line(line):
    UNFINISHED_GATES.append(Gate(line))


def main():
    with open('input7.txt') as infile:
        line = infile.readline()
        while line:
            process_line(line)
            line = infile.readline()
    print "finished setup. Start computation."
    while len(UNFINISHED_GATES) > 0:
        print str(len(UNFINISHED_GATES)) + " unfinished gates."
        # print WIRES
        for gate in UNFINISHED_GATES:
            if gate.can_compute():
                gate.execute()
                WIRES[gate.get_wirename()] = gate.get_output()
                UNFINISHED_GATES.remove(gate)
                print "gate finished: " + gate.get_wirename() + " : " + str(gate.get_output())
    print value_of('a')

if __name__ == '__main__':
    main()
