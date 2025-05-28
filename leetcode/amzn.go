Q1 - (All test cases passed)


The management team at one of Amazon's warehouse facilities is planning to install a smart lighting system.
A total of n smart bulbs will be installed along a straight line. A smart bulb automatically turns OFF if the sum of the brightness of the bulbs in front of it is greater than its own brightness. The bulbs can be rearranged in any order. The team wants to find out the minimum number of bulbs that will turn OFF across all possible permutations of the sequence of bulbs.
Given an array, brightness, that represents the brightness of the bulbs placed along a straight line, find the minimum number of bulbs that will turn OFF if the bulbs can be rearranged in any order.
Note: All bulbs are turned ON initially. They turn OFF (in order, from left to right) based on the condition mentioned above.
Example
Suppose, n = 5, brightness = [2, 1, 3, 4, 3]
Some of the possible arrangements of bulbs, their final states, and the corresponding number of OFF bulbs are as follows -


Arrangement | Final State | OFF Count


Note: T = ON state, F = OFF state


One of the optimal arrangements is [1, 2, 3, 4, 3]. The minimum number of OFF bulbs in the final state is 2.


    public static int getMinimumOffBulbs(List<Integer> brightness) {
        Collections.sort(brightness);
        int runningSum = 0;
        int offCount = 0;
        
        for (int b : brightness) {
            if (runningSum > b) {
                offCount++;
            } else {
                runningSum += b;
            }
        }
        
        return offCount;
    }
Q2 - (All test cases passed)


Amazon has multiple delivery centers and delivery warehouses all over the world! The world is represented by a number line from - 10 ^ 9 to 10 ^ 9 There are n delivery centers, the one at location center[i]. A location x is called a suitable location for a warehouse if it is possible to bring all the products to that point by traveling a distance of no more than d. At any one time, products can bejorought from one delivery center and placed at point x. Given the positions of n delivery centers, calculate the number of suitable locations in the world. That is, calculate the number of points x on the number line ( - 10 ^ 9 <=*<=10^ 9 ) where the travel distance required to bring all the products to that point is less than or equal to d.


Note: The distance between point x and center[i] is | x- center[i] |, their absolute difference.


Given n = 3 center =[ - 2, 1 , 0 ], d = 8


At x = -3: Total distance = 16 > d, not suitable
At x = 0: Total distance = 6 ≤ d, suitable
At x = -1: Total distance = 8 ≤ d, suitable
At x = 1: Total distance = 8 ≤ d, suitable


The function should return 3, as there are three suitable locations (-1, 0, 1).


Constraints
1 ≤ n ≤ 10^5 (where n is the number of delivery centers)
-10^9 ≤ center[i] ≤ 10^9
0 ≤ d ≤ 10^15


   public static int suitableLocations(List<Integer> center, long d) {
        Collections.sort(center);
        int median = center.get(center.size() / 2);

        if (!totalDistGood(center, median, d)) {
            return 0;
        }

        int min = findMin(center, median, d);
        int max = findMax(center, median, d);

        return max - min + 1;
    }

    private static int findMax(List<Integer> center, int x, long d) {
        int l = x;
        int r = 1000000000;
        int re = x;
        while (l <= r) {
            int m = l + (r - l) / 2;
            if (totalDistGood(center, m, d)) {
                re = m;
                l = m + 1;
            } else {
                r = m - 1;
            }
        }
        return re;
    }

    private static int findMin(List<Integer> center, int x, long d) {
        int l = -1000000000;
        int r = x;
        int re = x;
        while (l <= r) {
            int m = l + (r - l) / 2;
            if (totalDistGood(center, m, d)) {
                re = m;
                r = m - 1;
            } else {
                l = m + 1;
            }
        }
        return re;
    }

    private static boolean totalDistGood(List<Integer> center, int x, long d) {
        long totalDistance = 0;
        for (int c : center) {
            totalDistance += 2L * Math.abs(c - x);
            if (totalDistance > d) {
                return false;
            }
        }
        return true;
    }