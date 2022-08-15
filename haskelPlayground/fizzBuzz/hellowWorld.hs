-- based on https://www.youtube.com/watch?v=02_H3LjqMr8
-- web ide https://app.codingrooms.com/


-- imports
import Data.List

------------------Data Types
-- Int -- bound
-- Integer -- unbound number
-- Float
-- Double
-- Bool True False
-- Char '
-- Tuple -- 
maxInt = maxBound :: Int

--constant
always5 :: Integer
always5 = 5


------------------Math
-- + * - /
modEx = mod 5 4
modEx2 = 5 `mod` 4 --infix operator
negNumEx = 5 + (-4)
sumOfNums = sum [1..1000]

--sqrt :: Floating a => a -> a


main :: IO ()
main = do
    putStrLn "Start"
    putStrLn ""
    putStrLn "End"
