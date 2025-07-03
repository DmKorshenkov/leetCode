/** 
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad 
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

func firstBadVersion(n int) int {
    lo , hi := 1, n

    for lo <= hi {
        v := lo + (hi - lo)/2

        if isBadVersion(v) {
            hi = v - 1
        } else {
            lo = v + 1
        }
    }
    return lo
}
