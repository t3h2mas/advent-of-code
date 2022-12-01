package xyz.t3h2mas.advent.one.domain;

public class FuelCounterUpper {
    private final FuelStrategy fuelStrategy;

    public FuelCounterUpper(FuelStrategy fuelStrategy) {
        this.fuelStrategy = fuelStrategy;
    }

    public int calculateFuelRequirements(SpaceCraft sc) {
        return this.fuelStrategy.fuelRequiredBy(sc);
    }
}
