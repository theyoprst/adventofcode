import Foundation

// MARK: - Public API

/// Run solutions interactively with input file selection menu.
public func runInteractively<T>(
    part1Solutions: [Solution<T>],
    part2Solutions: [Solution<T>],
    bundle: Bundle
) {
    let lines = selectAndLoadInput(bundle: bundle)

    for solution in part1Solutions {
        print("Part 1 (\(solution.name)):", solution.solve(lines))
    }
    for solution in part2Solutions {
        print("Part 2 (\(solution.name)):", solution.solve(lines))
    }
}

// MARK: - Private Helpers

private func selectAndLoadInput(bundle: Bundle) -> [String] {
    let inputFiles = discoverInputFiles(bundle: bundle)

    guard !inputFiles.isEmpty else {
        print("Error: No input*.txt files found in resources")
        print("Make sure input files are added to Package.swift resources")
        exit(1)
    }

    print("Available input files:")
    for (index, file) in inputFiles.enumerated() {
        print("  \(index + 1). \(file)")
    }
    print()
    print("Select input file (1-\(inputFiles.count)): ", terminator: "")

    guard let input = readLine(),
          let selection = Int(input),
          selection >= 1 && selection <= inputFiles.count else {
        print("Error: Invalid selection")
        exit(1)
    }

    let selectedFile = inputFiles[selection - 1]
    guard let resourcePath = bundle.resourcePath else {
        print("Error: Could not access bundle resources")
        exit(1)
    }

    let filePath = (resourcePath as NSString).appendingPathComponent(selectedFile)
    guard let data = try? String(contentsOfFile: filePath) else {
        print("Error: Could not read file: \(selectedFile)")
        exit(1)
    }

    var lines = data.split(separator: "\n", omittingEmptySubsequences: false)
        .map(String.init)
    while lines.last?.isEmpty == true {
        lines.removeLast()
    }

    print("\nRunning with: \(selectedFile)")
    print("---")

    return lines
}

private func discoverInputFiles(bundle: Bundle) -> [String] {
    guard let resourcePath = bundle.resourcePath else {
        return []
    }

    let fileManager = FileManager.default
    guard let contents = try? fileManager.contentsOfDirectory(atPath: resourcePath) else {
        return []
    }

    let inputFiles = contents.filter { $0.hasPrefix("input") && $0.hasSuffix(".txt") }

    let mainInput = inputFiles.filter { $0 == "input.txt" }
    let exampleInputs = inputFiles.filter { $0 != "input.txt" }.sorted()

    return exampleInputs + mainInput
}
