package xyz.t3h2mas.advent.three;

import org.junit.Test;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

import static org.junit.Assert.*;

public class SchematicReaderShould {
    @Test
    public void
    translate_input_into_wire_paths() {
        String input = "R75,D30,R83,U83,L12,D49,R71,U7,L72\n" +
                "U62,R66,U55,R34,D71,R55,D58,R83";

        List<String>[] expected = new ArrayList[2];
        expected[0] = new ArrayList<>(Arrays.asList(
                "R75", "D30", "R83", "U83", "L12", "D49", "R71", "U7", "L72"
        ));

        expected[1] = new ArrayList<>(Arrays.asList(
                "U62", "R66", "U55", "R34", "D71", "R55", "D58", "R83"
        ));

        assertArrayEquals(expected, SchematicReader.readFrom(input));
    }

}