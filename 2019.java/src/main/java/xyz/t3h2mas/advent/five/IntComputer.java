package xyz.t3h2mas.advent.five;

public class IntComputer {
    private final int OP_CODE_ADD = 1;
    private final int OP_CODE_MULTIPLY = 2;
    private final int OP_CODE_STORE = 3;
    private final int OP_CODE_PRINT = 4;
    private final int OP_CODE_HALT = 99;

    private IntComputerIO io;

    public IntComputer() {
        this.io = new IntComputerIO(System.in, System.out);
    }

    public void compute(int[] program) throws HaltSignaledException, IllegalStateException {
        int i = 0;
        while (i < program.length) {
            OpCode opCode = OpCode.fromInt(program[i]);

            switch(opCode.getCode()) {
                case OP_CODE_ADD:
                    handleAddition(program, opCode, i);
                    i += 4;
                    break;
                case OP_CODE_MULTIPLY:
                    handleMultiplication(program, opCode, i);
                    i += 4;
                    break;
                case OP_CODE_STORE:
                    handleStore(program, opCode, i);
                    i += 2;
                    break;
                case OP_CODE_PRINT:
                    handlePrint(program, opCode, i);
                    i += 2;
                case OP_CODE_HALT:
                    handleHalt();
                    break;
                default:
                    throw new IllegalStateException("Unexpected value: " + program[i]);
            }

        }
    }

    private void handleAddition(int[] program, OpCode opCode, int currentPosition) {
        int firstArgAddr = program[currentPosition + 1];
        int secondArgAddr = program[currentPosition + 2];
        int resultAddr = program[currentPosition + 3];

        int firstValue;
        if (opCode.parameterModeFor(0) == ParameterMode.POSITION) {
           firstValue = program[firstArgAddr];
        } else {
            firstValue = firstArgAddr;
        }

        int secondValue;
        if (opCode.parameterModeFor(1) == ParameterMode.POSITION) {
            secondValue = program[secondArgAddr];
        } else {
            secondValue = secondArgAddr;
        }
        program[resultAddr] = firstValue + secondValue;
    }
    private void handleMultiplication(int[] program, OpCode opCode, int currentPosition) {
        int firstArgAddr = program[currentPosition + 1];
        int secondArgAddr = program[currentPosition + 2];
        int resultAddr = program[currentPosition + 3];


        int firstValue;
        if (opCode.parameterModeFor(0) == ParameterMode.POSITION) {
            firstValue = program[firstArgAddr];
        } else {
            firstValue = firstArgAddr;
        }

        int secondValue;
        if (opCode.parameterModeFor(1) == ParameterMode.POSITION) {
            secondValue = program[secondArgAddr];
        } else {
            secondValue = secondArgAddr;
        }

        program[resultAddr] = firstValue * secondValue;
    }
    private void handleStore(int[] program, OpCode opCode, int currentPosition) {
        int storeAtAddr = program[currentPosition + 1];
        int valueToStore = this.io.readInt();
        program[storeAtAddr] = valueToStore;
    }
    private void handlePrint(int[] program, OpCode opCode, int currentPosition) {
        int firstArgAddr = program[currentPosition + 1];
        int firstValue;
        if (opCode.parameterModeFor(0) == ParameterMode.POSITION) {
            firstValue = program[firstArgAddr];
        } else {
            firstValue = firstArgAddr;
        }
        System.out.println("Printing");
        this.io.printInt(firstValue);
    }
    private void handleHalt() throws HaltSignaledException {
        throw new HaltSignaledException("Halt signaled.");
    }
}
