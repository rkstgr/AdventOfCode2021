import time
from itertools import cycle

import numpy as np


class Line:

    def __init__(self, start, end):
        self.start = np.array([start[0], start[1]])
        self.end = np.array([end[0], end[1]])

    def onGrid(self) -> bool:
        """Checks if line is vertical or horizontal

        :return: bool
        """
        if self.start[0] == self.end[0] or self.start[1] == self.end[1]:
            return True
        else:
            return False

    def __repr__(self):
        return f"{self.start} -> {self.end}"


def parseLineInput(lineInput: str) -> Line:
    startInput, endInput = lineInput.split(" -> ")
    startPoints = [int(x) for x in startInput.split(",")]
    endPoints = [int(x) for x in endInput.split(",")]
    line = Line(startPoints, endPoints)
    return line


def getBounds(lines: [Line]):
    points = []
    for line in lines:
        points.append(line.start)
        points.append(line.end)
    pointMatrix = np.stack(points)
    xMax = pointMatrix[:, 0].max()
    yMax = pointMatrix[:, 1].max()
    return xMax, yMax


def part1(lines):
    x, y = getBounds(lines)
    ground = np.zeros((x + 1, y + 1))

    for line in lines:
        if line.onGrid():
            xs, xe = line.start[0], line.end[0]
            ys, ye = line.start[1], line.end[1]
            xPoints = np.arange(xs, xe + 1) if xe > xs else np.arange(xe, xs + 1)
            yPoints = np.arange(ys, ye + 1) if ye > ys else np.arange(ye, ys + 1)
            for i in range(max(len(xPoints), len(yPoints))):
                x = xPoints[i % len(xPoints)]
                y = yPoints[i % len(yPoints)]
                ground[x, y] += 1

    return np.sum(ground >= 2)


def part2(lines):
    x, y = getBounds(lines)
    ground = np.zeros((x + 1, y + 1))

    for line in lines:
        xs, xe = line.start[0], line.end[0]
        ys, ye = line.start[1], line.end[1]
        xPoints = np.arange(xs, xe + 1) if xe > xs else np.arange(xs, xe - 1, -1)
        yPoints = np.arange(ys, ye + 1) if ye > ys else np.arange(ys, ye - 1, -1)
        for i in range(max(len(xPoints), len(yPoints))):
            x = xPoints[i % len(xPoints)]
            y = yPoints[i % len(yPoints)]
            ground[x, y] += 1

    return np.sum(ground >= 2)


if __name__ == '__main__':
    start = time.time()
    # parse input
    with open("input.txt") as f:
        lines = [parseLineInput(x) for x in f.readlines()]

    sol1 = part1(lines)
    sol2 = part2(lines)

    print(sol1)
    print(sol2)

    print(f"time: {(time.time() - start).__round__(6) * 1000 } ms")
