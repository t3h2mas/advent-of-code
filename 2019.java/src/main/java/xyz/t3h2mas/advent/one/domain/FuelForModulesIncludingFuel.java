package xyz.t3h2mas.advent.one.domain;

public class FuelForModulesIncludingFuel implements FuelStrategy {
    private int calculateRequiredFuelFor(int mass) {
        return (int) Math.floor((float) mass / 3.0) - 2;
    }

    private int calculateCompositeFuelRequirementsFor(int mass) {
        int total = 0;
        int currentMass = mass;
        while (true) {
            currentMass = this.calculateRequiredFuelFor(currentMass);
            if (currentMass < 1) {
                break;
            }
            total += currentMass;
        }
        return total;
    }

    @Override
    public int fuelRequiredBy(SpaceCraft sc) {
        return sc.getModules()
                .stream()
                .map(m -> this.calculateCompositeFuelRequirementsFor(m.getMass()))
                .reduce(0, Integer::sum);
    }

}
