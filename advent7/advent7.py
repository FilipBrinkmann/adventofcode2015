"""Solver for advent of code no.7"""


class Gate(object):
    """Represents a gate object."""

    ASSIGN = '->'

    def __init__(self, description):
        super(Gate, self).__init__()
        self.wire = self._extractWire(description)
        self.gate_type = self._extract_gate_type(description)
        self.inputs = self._extract_inputs(description)
        self.output = None
        self.ready_to_compute = self.can_compute()
        WIRES[self.wire] = self.output
        for inp in self.inputs:
            if inp not in WIRES.keys() and not self._is_number(inp):
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
        for i in range(len(self.inputs)):
            if self.inputs[i] in WIRES.keys() and WIRES[self.inputs[i]] is not None:
                self.inputs[i] = WIRES[self.inputs[i]].get_output()

    def _convert_inputs_to_ints(self):
        for i in range(len(self.inputs)):
            if self._is_number(self.inputs[i]):
                self.inputs[i] = int(self.inputs[i])

    def can_compute(self):
        for inp in self.inputs:
            if not self._is_number(inp):
                return False
        return True

    def execute(self):
        if self.output is not None:
            raise Exception("you must not call execute() more than once")

        if not self.can_compute():
            self._update_inputs()
            for inputwire in self.inputs:
                if not self._is_number(inputwire):
                    responsible_gate = WIRES[inputwire]
                    if responsible_gate is None:
                        print self.inputs
                        raise Exception('no responsible gate for wire '+inputwire)
                    responsible_gate.execute()
                    self.inputs[self.inputs.index(inputwire)] = responsible_gate.get_output()

        self._convert_inputs_to_ints()
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
        if self.output is None:
            self.execute()
        return self.output


WIRES = {}


def process_line(line):
    return Gate(line)


def main():
    with open('input7.txt') as infile:
        line = infile.readline()
        while line:
            gate = process_line(line)
            WIRES[gate.get_wirename()] = gate
            line = infile.readline()
    print "Finished setup. Starting computation."
    a_gate = WIRES['a']
    a_gate.execute()
    print "value of wire '" + a_gate.get_wirename() + "' : " + str(a_gate.get_output())

if __name__ == '__main__':
    main()
