package com.architecture.clean.archunit;

import com.tngtech.archunit.core.domain.JavaClasses;
import com.tngtech.archunit.core.importer.ClassFileImporter;

import java.util.List;

import static com.tngtech.archunit.base.DescribedPredicate.greaterThanOrEqualTo;
import static com.tngtech.archunit.lang.conditions.ArchConditions.containNumberOfElements;
import static com.tngtech.archunit.lang.syntax.ArchRuleDefinition.classes;
import static com.tngtech.archunit.lang.syntax.ArchRuleDefinition.noClasses;

abstract class ArchitectureElement {

    final String basePackage;

    public ArchitectureElement(String basePackage) {
        this.basePackage = basePackage;
    }

    String fullQualifiedPackage(String relativePackage) {
        return this.basePackage + "." + relativePackage;
    }

    static void denyDependency(String fromPackageName, String toPackageName, JavaClasses classes) {
        noClasses()
                .that()
                .resideInAPackage("io.reflectoring.reviewapp.domain..")
                .should()
                .dependOnClassesThat()
                .resideInAnyPackage("io.reflectoring.reviewapp.application..")
                .check(classes);
    }

    static void denyAnyDependency(
            List<String> fromPackages, List<String> toPackages, JavaClasses classes) {
        for (String fromPackage : fromPackages) {
            for (String toPackage : toPackages) {
                noClasses()
                        .that()
                        .resideInAPackage(matchAllClassesInPackage(fromPackage))
                        .should()
                        .dependOnClassesThat()
                        .resideInAnyPackage(matchAllClassesInPackage(toPackage))
                        .check(classes);
            }
        }
    }

    static String matchAllClassesInPackage(String packageName) {
        return packageName + "..";
    }

    void denyEmptyPackage(String packageName) {
        classes()
                .that()
                .resideInAPackage(matchAllClassesInPackage(packageName))
                .should(containNumberOfElements(greaterThanOrEqualTo(1)))
                .check(classesInPackage(packageName));
    }

    private JavaClasses classesInPackage(String packageName) {
        return new ClassFileImporter().importPackages(packageName);
    }

    void denyEmptyPackages(List<String> packages) {
        for (String packageName : packages) {
            denyEmptyPackage(packageName);
        }
    }
}

