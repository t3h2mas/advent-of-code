package xyz.t3h2mas.advent.three.helpers;

import org.junit.Test;
import xyz.t3h2mas.advent.three.Point;


import java.util.HashSet;
import java.util.Set;

import static org.junit.Assert.*;

public class SetHelperShould {
    @Test
    public void
    find_intersection_of_two_point_sets() {
        Set<Point> s1 = new HashSet<>();
        s1.add(new Point(42, 42));
        s1.add(new Point(41, 42));

        Set<Point> s2 = new HashSet<>();
        s2.add(new Point(42, 42));
        s2.add(new Point(42, 41));

        Set<Point> intersections = SetHelper.intersect(s1, s2);

        Set<Point> expected = new HashSet<>();
        expected.add(new Point(42, 42));
        assertEquals(intersections, expected);
    }
}