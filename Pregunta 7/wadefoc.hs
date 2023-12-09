import System.Environment
import Data.List

fibonacciChola :: Integer -> Integer
fibonacciChola = fibAux 1 0 where
    fibAux :: Integer -> Integer -> Integer -> Integer
    fibAux x y k = if k == 0 then x else fibAux (x+y) x (k-1)

wadefoc :: Integer -> Integer
wadefoc n = fibonacciChola chorizo where
    naranyana = fromIntegral ((n + 1) * n * n * (n - 1)) / 12
    chorizo = fibonacciChola (floor (logBase 2 naranyana) + 1)

main :: IO ()
main = do
    [arg] <- getArgs
    let n = read arg :: Integer
    print $ wadefoc n
