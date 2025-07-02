/** 
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 * func guess(num int) int;
 */

func guessNumber(n int) int {
    low, hight := 1, n

    for low <= hight {
        pick := low + (hight-low)/2
        
        if ok := guess(pick); ok == 0 {
            return pick
        } else if ok == -1 {
            hight = pick -1
        } else {
            low = pick + 1
        }
    }
    return 0
}
