// swift-tools-version: 6.0
import PackageDescription

let package = Package(
    name: "AdventOfCode2025",
    platforms: [
        .macOS(.v13),
        .iOS(.v16),
        .tvOS(.v16),
        .watchOS(.v9),
    ],
    products: [
        .executable(name: "day01", targets: ["Day01"]),
        .executable(name: "day02", targets: ["Day02"]),
        .executable(name: "day03", targets: ["Day03"]),
        .executable(name: "day04", targets: ["Day04"]),
        .executable(name: "day05", targets: ["Day05"]),
        .executable(name: "day06", targets: ["Day06"]),
    ],
    dependencies: [
        .package(url: "https://github.com/jpsim/Yams.git", from: "5.0.0"),
    ],
    targets: [
        .target(
            name: "AOCUtilities",
            path: "AOCUtilities",
        ),
        .target(
            name: "AOCTestSupport",
            dependencies: ["Yams", "AOCUtilities"],
            path: "AOCTestSupport",
        ),
        .executableTarget(
            name: "Day01",
            dependencies: ["AOCUtilities"],
            path: "01",
            exclude: ["SolutionTests.swift", "tests.yaml", "part1.html", "part1.md", "part2.html", "part2.md"],
            sources: ["Solution.swift"],
            resources: [
                .copy("input.txt"),
                .copy("input_ex1.txt"),
            ]
        ),
        .executableTarget(
            name: "Day02",
            dependencies: ["AOCUtilities"],
            path: "02",
            exclude: ["SolutionTests.swift", "tests.yaml", "part1.html", "part1.md", "part2.html", "part2.md"],
            sources: ["Solution.swift"],
            resources: [
                .copy("input.txt"),
                .copy("input_ex1.txt"),
            ]
        ),
        .executableTarget(
            name: "Day03",
            dependencies: ["AOCUtilities"],
            path: "03",
            exclude: ["SolutionTests.swift", "tests.yaml", "part1.html", "part1.md", "part2.html", "part2.md"],
            sources: ["Solution.swift"],
            resources: [
                .copy("input.txt"),
                .copy("input_ex1.txt"),
            ]
        ),
        .executableTarget(
            name: "Day04",
            dependencies: ["AOCUtilities"],
            path: "04",
            exclude: ["SolutionTests.swift", "tests.yaml", "part1.html", "part1.md", "part2.html", "part2.md"],
            sources: ["Solution.swift"],
            resources: [
                .copy("input.txt"),
                .copy("input_ex1.txt"),
            ]
        ),
        .executableTarget(
            name: "Day05",
            dependencies: ["AOCUtilities"],
            path: "05",
            exclude: ["SolutionTests.swift", "tests.yaml", "part1.html", "part1.md", "part2.html", "part2.md"],
            sources: ["Solution.swift"],
            resources: [
                .copy("input.txt"),
                .copy("input_ex1.txt"),
            ]
        ),
        .executableTarget(
            name: "Day06",
            dependencies: ["AOCUtilities"],
            path: "06",
            exclude: ["SolutionTests.swift", "tests.yaml", "part1.html", "part1.md", "part2.html", "part2.md"],
            sources: ["Solution.swift"],
            resources: [
                .copy("input.txt"),
                .copy("input_ex1.txt"),
            ]
        ),
        .testTarget(
            name: "Day01Tests",
            dependencies: ["Day01", "AOCTestSupport"],
            path: "01",
            exclude: ["Solution.swift", "part1.html", "part1.md", "part2.html", "part2.md"],
            sources: ["SolutionTests.swift"],
            resources: [
                .copy("tests.yaml"),
                .copy("input.txt"),
                .copy("input_ex1.txt"),
            ]
        ),
        .testTarget(
            name: "Day02Tests",
            dependencies: ["Day02", "AOCTestSupport"],
            path: "02",
            exclude: ["Solution.swift", "part1.html", "part1.md", "part2.html", "part2.md"],
            sources: ["SolutionTests.swift"],
            resources: [
                .copy("tests.yaml"),
                .copy("input.txt"),
                .copy("input_ex1.txt"),
            ]
        ),
        .testTarget(
            name: "Day03Tests",
            dependencies: ["Day03", "AOCTestSupport"],
            path: "03",
            exclude: ["Solution.swift", "part1.html", "part1.md", "part2.html", "part2.md"],
            sources: ["SolutionTests.swift"],
            resources: [
                .copy("tests.yaml"),
                .copy("input.txt"),
                .copy("input_ex1.txt"),
            ]
        ),
        .testTarget(
            name: "Day04Tests",
            dependencies: ["Day04", "AOCTestSupport"],
            path: "04",
            exclude: ["Solution.swift", "part1.html", "part1.md", "part2.html", "part2.md"],
            sources: ["SolutionTests.swift"],
            resources: [
                .copy("tests.yaml"),
                .copy("input.txt"),
                .copy("input_ex1.txt"),
            ]
        ),
        .testTarget(
            name: "Day05Tests",
            dependencies: ["Day05", "AOCTestSupport"],
            path: "05",
            exclude: ["Solution.swift", "part1.html", "part1.md", "part2.html", "part2.md"],
            sources: ["SolutionTests.swift"],
            resources: [
                .copy("tests.yaml"),
                .copy("input.txt"),
                .copy("input_ex1.txt"),
            ]
        ),
        .testTarget(
            name: "Day06Tests",
            dependencies: ["Day06", "AOCTestSupport"],
            path: "06",
            exclude: ["Solution.swift", "part1.html", "part1.md", "part2.html", "part2.md"],
            sources: ["SolutionTests.swift"],
            resources: [
                .copy("tests.yaml"),
                .copy("input.txt"),
                .copy("input_ex1.txt"),
            ]
        ),
    ]
)
