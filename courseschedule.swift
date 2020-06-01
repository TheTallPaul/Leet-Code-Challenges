/*
210. Course Schedule II

There are a total of n courses you have to take, labeled from 0 to n-1.
Some courses may have prerequisites, for example to take course 0 you have to first take course 1, which is expressed as a pair: [0,1]
Given the total number of courses and a list of prerequisite pairs, return the ordering of courses you should take to finish all courses.
There may be multiple correct orders, you just need to return one of them. If it is impossible to finish all courses, return an empty array.

https://leetcode.com/problems/course-schedule-ii/
*/

class Course {
    var id: Int
    var unlocks: [Int] = []
    var reqCount: Int
    
    init (id: Int) {
        self.id = id
        reqCount = 0
    }
    
    func addUnlock(req: Int) {
        unlocks.append(req)
    }
    
    func incReqCount() {
        reqCount += 1
    }
    
    func decReqCount() {
        reqCount -= 1
    }
}


class Solution {

    func findOrder(_ numCourses: Int, _ prerequisites: [[Int]]) -> [Int] {
        var courseList = [Course]()
        var freeCourses = [Course]()
        var courseOrder = [Int]()
        
        // Create list of possible courses
        for i in 0...numCourses - 1 {
            courseList.append(Course(id: i))
        }
        
        // Adds unlocks and requirement counts to the courses
        for requirement in prerequisites {
            courseList[requirement[1]].addUnlock(req: requirement[0])
            courseList[requirement[0]].incReqCount()
        }
        
        // Create a list of courses without requirements
        for (i, course) in courseList.enumerated() {
            if courseList[i].reqCount == 0 {
                freeCourses.append(course)
            }
        }
        
        while freeCourses.count != 0 && courseOrder.count < numCourses {
            courseOrder.append(freeCourses[0].id)
            
            // Update the reqCount for the courses it unlocks
            for unlockedClass in freeCourses[0].unlocks {
                courseList[unlockedClass].decReqCount()
                
                // Add courses that are now unlocked
                if courseList[unlockedClass].reqCount == 0 {
                    freeCourses.append(courseList[unlockedClass])
                }
            }
            
            // Pop the course
            freeCourses.remove(at: 0)
        }
        
        if courseOrder.count == numCourses {
            return courseOrder
        }
        
        return [Int]()
    }
}
