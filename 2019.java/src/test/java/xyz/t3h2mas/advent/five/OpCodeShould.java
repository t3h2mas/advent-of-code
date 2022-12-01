package xyz.t3h2mas.advent.five;

import org.junit.Test;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

import static org.junit.Assert.*;

public class OpCodeShould {

    @Test
    public void extract_the_op_code_from_an_int() {
        int OP_CODE_RAW = 1002;

        OpCode actual = OpCode.fromInt(OP_CODE_RAW);

        assertEquals(2, actual.getCode());
    }

    @Test
    public void extract_the_parameter_modes_from_int() {
        int OP_CODE_RAW = 1002;

        List<ParameterMode> expected = new ArrayList<>(Arrays.asList(ParameterMode.POSITION, ParameterMode.IMMEDIATE));

        OpCode actual = OpCode.fromInt(OP_CODE_RAW);
        assertEquals(expected, actual.getParameterModes());
    }

    @Test public void
    get_a_parameter_mode_by_index() {
        int OP_CODE_RAW = 111002;

        OpCode actual = OpCode.fromInt(OP_CODE_RAW);
        assertEquals(ParameterMode.POSITION, actual.parameterModeFor(0));
        assertEquals(ParameterMode.IMMEDIATE, actual.parameterModeFor(1));
        assertEquals(ParameterMode.IMMEDIATE, actual.parameterModeFor(2));
        assertEquals(ParameterMode.IMMEDIATE, actual.parameterModeFor(3));
    }

    @Test public void
    get_default_parameter_mode_for_non_existent_index() {
        int OP_CODE_RAW = 1102;

        OpCode actual = OpCode.fromInt(OP_CODE_RAW);
        assertEquals(ParameterMode.IMMEDIATE, actual.parameterModeFor(0));
        assertEquals(ParameterMode.IMMEDIATE, actual.parameterModeFor(1));
        assertEquals(ParameterMode.POSITION, actual.parameterModeFor(2));
        assertEquals(ParameterMode.POSITION, actual.parameterModeFor(3));
    }
}