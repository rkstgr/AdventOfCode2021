import numpy as np
import torch
from torch.autograd import Variable


def triangle_distance(x, y):
    l1 = torch.abs(x - y)
    return l1 * (l1 + 1) / 2


def total_distance(nums, x):
    return triangle_distance(nums, x).sum()


def part2(nums: np.ndarray):
    positions = torch.from_numpy(nums).double()

    x = Variable(torch.mean(positions).type(torch.DoubleTensor), requires_grad=True)

    # we want to minimize our distance = loss
    loss = total_distance(positions, x)
    print("initial distance:", loss)

    lr = 1e-3
    for i in range(4):
        loss = total_distance(positions, x)
        print(f"[{i}] x", x)
        print(f"[{i}] distance", loss)

        loss.backward()

        x.data -= lr * x.grad.data
        x.grad.data.zero_()

    x = x.detach().round()
    return total_distance(positions, x)


if __name__ == '__main__':
    with open("input.txt") as f:
        nums = [int(x) for x in f.readline().split(",")]
    nums = np.array(nums)

    # Part 1
    # The median minimizes the sum of l1 distances
    x = np.median(nums)
    sol1 = np.abs(x - nums).sum()

    # Part 2
    # Minimize distance with SGD starting from the mean
    sol2 = part2(nums)

    print("---\nSolutions")
    print("Part 1:", sol1)
    print("Part 2:", sol2.numpy())
