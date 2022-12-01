package xyz.t3h2mas.advent.five;


import java.util.ArrayList;
import java.util.List;

public class OpCode {
    private final int code;
    private final List<ParameterMode> parameterModes;

    private final ParameterMode DEFAULT_PARAM_MODE = ParameterMode.POSITION;

    private OpCode(int code, List<ParameterMode> parameterModes) {
        this.code = code;
        this.parameterModes = parameterModes;
    }

    public static OpCode fromInt(int n) {
        List<ParameterMode> modes = new ArrayList<>();

        String s = Integer.toString(n);

        int code;
        if (s.length() <= 2) {
            code = Integer.parseInt(s);
        } else {
            code = Integer.parseInt(s.substring(s.length() - 2));
        }

        if (s.length() > 2) {
            String[] modeChunks = s.substring(0, s.length() - 2).split("");

            // The modes are stored from right to left, unpack starting at the end of the list
            for (int i = modeChunks.length - 1; i >= 0; i--) {
                switch (modeChunks[i]) {
                    case "0":
                        modes.add(ParameterMode.POSITION);
                        break;
                    case "1":
                        modes.add(ParameterMode.IMMEDIATE);
                        break;
                    default:
                        System.err.println("Warning: Unknown parameter mode: " + modeChunks[i]);
                }
            }

        }
        return new OpCode(code, modes);
    }

    public int getCode() {
        return code;
    }

    public List<ParameterMode> getParameterModes() {
        return parameterModes;
    }

    public ParameterMode parameterModeFor(int idx) {
        try {
            return this.parameterModes.get(idx);
        } catch (IndexOutOfBoundsException e) {
            return this.DEFAULT_PARAM_MODE;
        }
    }
}
