plugins {
    id("org.jetbrains.kotlin.jvm") version "1.9.0"

    application
}

repositories {
    mavenCentral()
}

dependencies {
    implementation(platform("org.jetbrains.kotlin:kotlin-bom"))
    implementation("org.jetbrains.kotlin:kotlin-stdlib-jdk8")
    implementation("com.google.guava:guava:30.1.1-jre")
}

kotlin {
    jvmToolchain(11)
}

application {
    mainClass.set("zenith.AppKt")
}

version = "0.0.0"

tasks {
    jar {
        manifest.attributes["Main-Class"] = application.mainClass

        val dependencies = configurations
            .runtimeClasspath
            .get()
            .map { zipTree(it) }

        from(dependencies)
        duplicatesStrategy = DuplicatesStrategy.EXCLUDE
    }
}
