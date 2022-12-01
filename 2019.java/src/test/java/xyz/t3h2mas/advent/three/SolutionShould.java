package xyz.t3h2mas.advent.three;

import org.junit.Test;

import static org.junit.Assert.*;

public class SolutionShould {

    @Test public void
    find_the_intersection_with_the_lowest_distance_from_the_central_port() throws Exception {
        String inputOne = "R75,D30,R83,U83,L12,D49,R71,U7,L72\n" +
                "U62,R66,U55,R34,D71,R55,D58,R83";
        assertEquals(159, Solution.wireDistanceClosestToCentralPort(inputOne));

        String inputTwo = "R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51\n" +
                "U98,R91,D20,R16,D67,R40,U7,R15,U6,R7";
        assertEquals(135, Solution.wireDistanceClosestToCentralPort(inputTwo));
    }

    @Test public void
    find_the_intersection_with_the_lowest_signal_delay() throws Exception {
        String inputOne = "R75,D30,R83,U83,L12,D49,R71,U7,L72\n" +
                "U62,R66,U55,R34,D71,R55,D58,R83";
        assertEquals(610, Solution.wireDistanceLowestSignalDelay(inputOne));

        String inputTwo = "R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51\n" +
                "U98,R91,D20,R16,D67,R40,U7,R15,U6,R7";
        assertEquals(410, Solution.wireDistanceLowestSignalDelay(inputTwo));
    }
}