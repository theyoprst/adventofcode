import Foundation
import AOCUtilities

/// Run solutions automatically with all discovered input files.
func runAllInputs(
    part1Solutions: [Solution],
    part2Solutions: [Solution],
    bundle: Bundle,
    daySubdirectory: String
) {
    let inputFiles = discoverInputFiles(bundle: bundle, subdirectory: daySubdirectory)

    guard !inputFiles.isEmpty else {
        print("Error: No input*.txt files found in resources/\(daySubdirectory)")
        exit(1)
    }

    for inputFile in inputFiles {
        print("=== Processing: \(inputFile) ===")

        let lines = loadInputLines(from: inputFile, bundle: bundle, subdirectory: daySubdirectory)

        for solution in part1Solutions {
            let result = solution.solve(lines)
            print("Part 1 (\(solution.name)): \(result)")
        }

        for solution in part2Solutions {
            let result = solution.solve(lines)
            print("Part 2 (\(solution.name)): \(result)")
        }

        print()
    }
}

// MARK: - Private Helpers

private func loadInputLines(
    from filename: String,
    bundle: Bundle,
    subdirectory: String
) -> [String] {
    guard let resourcePath = bundle.resourcePath else {
        print("Error: Could not access bundle resources")
        exit(1)
    }

    let filePath = (resourcePath as NSString).appendingPathComponent(subdirectory)
    let fullPath = (filePath as NSString).appendingPathComponent(filename)
    guard let data = try? String(contentsOfFile: fullPath) else {
        print("Error: Could not read file: \(subdirectory)/\(filename)")
        exit(1)
    }

    var lines = data.split(separator: "\n", omittingEmptySubsequences: false)
        .map(String.init)
    while lines.last?.isEmpty == true {
        lines.removeLast()
    }

    return lines
}

private func discoverInputFiles(bundle: Bundle, subdirectory: String) -> [String] {
    guard let resourcePath = bundle.resourcePath else {
        return []
    }

    let dayPath = (resourcePath as NSString).appendingPathComponent(subdirectory)
    let fileManager = FileManager.default
    guard let contents = try? fileManager.contentsOfDirectory(atPath: dayPath) else {
        return []
    }

    let inputFiles = contents.filter { $0.hasPrefix("input") && $0.hasSuffix(".txt") }

    let mainInput = inputFiles.filter { $0 == "input.txt" }
    let exampleInputs = inputFiles.filter { $0 != "input.txt" }.sorted()

    return exampleInputs + mainInput
}
