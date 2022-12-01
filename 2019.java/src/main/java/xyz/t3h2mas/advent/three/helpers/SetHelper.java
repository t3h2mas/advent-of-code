package xyz.t3h2mas.advent.three.helpers;

import java.util.HashSet;
import java.util.Set;

public class SetHelper{
    public static <T> Set<T> intersect(Set<T> s1, Set<T> s2) {
        Set<T> intersections = new HashSet<T>(s1);
        intersections.retainAll(s2);
        return intersections;
    }

    public static <T> T get(Set<T> all, T target) {
        for (T cur: all) {
            if (cur.equals(target)) {
                return cur;
            }
        }
        return null;
    }
}

