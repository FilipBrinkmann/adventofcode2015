"""advent of code 2015 no. 6"""

class Matrix(object):
    """represents the grid"""
    def __init__(self, width, height):
        super(Matrix, self).__init__()
        self.width = width
        self.height = height
        self.array = [[0 for x in range(width)] for x in range(height)]

    def _toggle(self, x, y):
        self.array[x][y] += 2
        # value = self.array[x][y]
        # if value == 0:
        #     self.array[x][y] = 1
        # else:
        #     self.array[x][y] = 0

    def _processArea(self, operation, x1, y1, x2, y2):
        assert x1 < x2 and y1 < y2
        for x in range(x1, x2+1):
            for y in range(y1, y2+1):
                if operation == 'toggle':
                    self._toggle(x, y)
                elif operation == 'on':
                    self._on(x, y)
                elif operation == 'off':
                    self._off(x, y)
                else:
                    raise AttributeError("illegal operation " + operation)

    def _on(self, x, y):
        self.array[x][y] += 1

    def _off(self, x, y):
        if self.array[x][y] > 0:
            self.array[x][y] -= 1

    def _isOn(self, x, y):
        return self.array[x][y] > 0

    def _getBrightness(self, x, y):
        return self.array[x][y]

    def getTotalBrighness(self):
        totalBrightness = 0
        for x in range(self.width):
            for y in range(self.height):
                totalBrightness += self.array[x][y]
        return totalBrightness

    def getLitLights(self):
        counter = 0
        totalChecked = 0
        for x in range(self.width):
            for y in range(self.height):
                totalChecked += 1
                if self._isOn(x, y):
                    counter += 1
        assert totalChecked == self.width * self.height
        return counter

    def execute(self, line):
        elements = line.split()
        if len(elements) < 4:
            raise Exception('invalid line')
        if elements[0] == 'turn':
            operation = elements[1]
            corner1 = elements[2]
            corner2 = elements[4]
            self._processArea(
                operation,
                int(corner1.split(',')[0]),
                int(corner1.split(',')[1]),
                int(corner2.split(',')[0]),
                int(corner2.split(',')[1]))
        elif elements[0] == 'toggle':
            operation = elements[0]
            corner1 = elements[1]
            corner2 = elements[3]
            self._processArea(
                operation,
                int(corner1.split(',')[0]),
                int(corner1.split(',')[1]),
                int(corner2.split(',')[0]),
                int(corner2.split(',')[1]))
        else:
            raise Exception("Illegal argument: "+elements[0])

        # if line.startswith('turn'):


def main():
    grid = Matrix(1000, 1000)
    assert grid.getLitLights() == 0
    with open('input6.txt') as infile:
        line = infile.readline()
        while line:
            grid.execute(line)
            line = infile.readline()
    print "lit lights: " + str(grid.getLitLights())
    print "brighness:" + str(grid.getTotalBrighness())

if __name__ == '__main__':
    main()
