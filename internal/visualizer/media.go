package visualizer

import (
	"dijkstra-visualizer/internal/algorithm"
	"dijkstra-visualizer/internal/graph"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
)

func GenerateFrames(g *graph.Graph, steps []algorithm.Step, outputDir string) error {
	for i, step := range steps {
		dotContent := GenerateDOT(g, step)
		dotFile := filepath.Join(outputDir, fmt.Sprintf("step_%03d.dot", i))
		pngFile := filepath.Join(outputDir, fmt.Sprintf("step_%03d.png", i))

		if err := os.WriteFile(dotFile, []byte(dotContent), 0644); err != nil {
			return fmt.Errorf("failed to write DOT file: %w", err)
		}

		cmd := exec.Command("dot", "-Tpng", dotFile, "-o", pngFile)
		if output, err := cmd.CombinedOutput(); err != nil {
			return fmt.Errorf("failed to convert DOT to PNG: %w\n%s", err, output)
		}
	}

	return nil
}

func CreateGIF(frameDir, outputPath string, delayMs int) error {
	files, err := filepath.Glob(filepath.Join(frameDir, "step_*.png"))
	if err != nil {
		return fmt.Errorf("failed to list PNG files: %w", err)
	}

	if len(files) == 0 {
		return fmt.Errorf("no PNG files found in %s", frameDir)
	}

	sort.Strings(files)

	delay := fmt.Sprintf("%d", delayMs/10)

	args := []string{
		"-delay", delay,
		"-loop", "0",
	}
	args = append(args, files...)
	args = append(args, outputPath)

	cmd := exec.Command("magick", args...)
	if _, err := cmd.CombinedOutput(); err != nil {
		cmd = exec.Command("convert", args...)
		if output, err := cmd.CombinedOutput(); err != nil {
			return fmt.Errorf("failed to create GIF: %w\nOutput: %s", err, output)
		}
	}

	return nil
}

func CleanupDotFiles(dir string) error {
	dotFiles, err := filepath.Glob(filepath.Join(dir, "*.dot"))
	if err != nil {
		return err
	}

	for _, file := range dotFiles {
		if err := os.Remove(file); err != nil {
			return err
		}
	}

	return nil
}
