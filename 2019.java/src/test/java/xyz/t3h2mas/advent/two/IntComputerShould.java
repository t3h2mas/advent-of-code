package xyz.t3h2mas.advent.two;

import org.junit.Test;

import static org.junit.Assert.*;

public class IntComputerShould {
    @Test public void
    return_the_expected_state_given_a_program() {
        IntComputer computer = new IntComputer();
        int[] program = {1,0,0,0,99};
        int[] expected = {2,0,0,0,99};
        try {
            computer.compute(program);
        } catch (HaltSignaledException e) {}

        assertArrayEquals(expected, program);
    }
}