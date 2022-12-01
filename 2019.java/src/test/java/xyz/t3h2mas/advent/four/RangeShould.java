package xyz.t3h2mas.advent.four;

import org.junit.Test;

import static org.junit.Assert.*;

public class RangeShould {
    @Test
    public void
    build_from_a_string() {
        String input = "123-456";

        Range result = Range.from(input);
        assertEquals(result.getStart(), 123);
        assertEquals(result.getStop(), 456);
    }
}