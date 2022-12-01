package xyz.t3h2mas.advent.five;

import org.junit.Test;

import java.util.Arrays;

import static org.junit.Assert.*;

public class IntComputerShould {
    @Test public void
    compute_a_program() {
        int[] input = {1002,4,3,4,33};
        int[] expected = {1002, 4, 3, 4, 99};

        IntComputer computer = new IntComputer();

        try {
            computer.compute(input);
        } catch (HaltSignaledException e) {
        }

        assertArrayEquals(expected, input);
    }
}