package xyz.t3h2mas.advent.one.domain;

import java.util.ArrayList;

public class SpaceCraft {
    private ArrayList<Module> modules;

    public SpaceCraft() {
        this.modules = new ArrayList<>();
    }

    public void addModule(Module m) {
        this.modules.add(m);
    }

    public ArrayList<Module> getModules() {
        return this.modules;
    }
}
