package ros

func export(path, file string, hide, verbose bool) Command {
	return Command{
		Path:    path,
		Command: "export",
		Filter: func() map[string]string {
			if file != "" {
				return map[string]string{
					"file": file,
				}
			}
			return nil
		}(),
		Flags: map[string]bool{
			"hide-sensitive": hide,
			"verbose":        verbose,
		},
	}
}

func (r *Ros) Export(path string, hide, verbose bool) ([]string, error) {
	return r.Run(export(path, "", hide, verbose))
}

func (r *Ros) ExportFile(path, file string, hide, verbose bool) ([]string, error) {
	return r.Run(export(path, file, hide, verbose))
}
