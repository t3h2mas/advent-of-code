package xyz.t3h2mas.advent.three;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

public class SchematicReader {
    public static List<String>[] readFrom(String input) {
        final String[] rawWirePaths = input.split("\n");

        final int WIRE_COUNT = rawWirePaths.length;

        List<String>[] wirePaths = new ArrayList[WIRE_COUNT];
        for (int i = 0; i < WIRE_COUNT; i++) {
            wirePaths[i] = new ArrayList<>(Arrays.asList(rawWirePaths[i].split(",")));
        }

        return wirePaths;
    }
}
