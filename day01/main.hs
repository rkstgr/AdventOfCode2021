module Main where
    
main = do
    txt <- readFile "input.txt"
    let nums = read <$> lines txt
    let n = countIncreases nums
    print n

    let windows = slidingWindows nums
    let k = countIncreases windows
    print k


countIncreases :: [Int] -> Int
countIncreases nums = sum [if x>y then 1 else 0 | (y, x) <- zip nums (tail nums)]

slidingWindows :: [Int] -> [Int]
slidingWindows x = zipWith3 (\a b c -> a+b+c) x (tail x) (tail $ tail x)