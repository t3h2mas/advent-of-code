package xyz.t3h2mas.advent.two;

public class IntComputer {
    private final int STEP_UNITS = 4;
    private final int OP_CODE_ADD = 1;
    private final int OP_CODE_MULTIPLY = 2;
    private final int OP_CODE_HALT = 99;


    public void compute(int[] program) throws HaltSignaledException, IllegalStateException {
        for(int i=0; i < program.length; i += STEP_UNITS) {
            switch(program[i]) {
                case OP_CODE_ADD:
                    handleAddition(program, i);
                    break;
                case OP_CODE_MULTIPLY:
                    handleMultiplication(program, i);
                    break;
                case OP_CODE_HALT:
                    handleHalt();
                    break;
                default:
                    throw new IllegalStateException("Unexpected value: " + program[i]);
            }
        }
    }

    private void handleAddition(int[] program, int currentPosition) {
        int firstArgAddr = program[currentPosition + 1];
        int secondArgAddr = program[currentPosition + 2];
        int resultAddr = program[currentPosition + 3];
        program[resultAddr] = program[firstArgAddr] + program[secondArgAddr];
    }
    private void handleMultiplication(int[] program, int currentPosition) {
        int firstArgAddr = program[currentPosition + 1];
        int secondArgAddr = program[currentPosition + 2];
        int resultAddr = program[currentPosition + 3];
        program[resultAddr] = program[firstArgAddr] * program[secondArgAddr];
    }
    private void handleHalt() throws HaltSignaledException {
        throw new HaltSignaledException("Halt signaled.");
    }
}
