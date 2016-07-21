package twentythousandtonnesofcrudeoil

import (
	"os"
	"reflect"
	"strings"

	"github.com/jessevdk/go-flags"
)

func TheEnvironmentIsPerfectlySafe(parser *flags.Parser, prefix string) {
	installEnv(parser, prefix)
}

func installEnv(parser *flags.Parser, prefix string) {
	eachOption(parser.Command, func(opt *flags.Option) {
		if len(opt.EnvDefaultKey) == 0 {
			opt.EnvDefaultKey = prefix + flagEnvName(opt.LongNameWithNamespace())

			kind := reflect.TypeOf(opt.Value()).Kind()
			if kind == reflect.Map || kind == reflect.Slice {
				opt.EnvDefaultDelim = ","
			}
		}
	})

	parser.CommandHandler = func(command flags.Commander, args []string) error {
		clearEnv(parser, prefix)

		return command.Execute(args)
	}
}

func clearEnv(parser *flags.Parser, prefix string) {
	eachOption(parser.Command, func(opt *flags.Option) {
		if strings.HasPrefix(opt.EnvDefaultKey, prefix) {
			os.Unsetenv(opt.EnvDefaultKey)
		}
	})
}

func flagEnvName(flag string) string {
	return strings.Replace(strings.ToUpper(flag), "-", "_", -1)
}

func eachOption(cmd *flags.Command, cb func(*flags.Option)) {
	eachOptionGroup(cmd.Group, cb)

	for _, group := range cmd.Groups() {
		eachOptionGroup(group, cb)
	}

	for _, sub := range cmd.Commands() {
		eachOption(sub, cb)
	}
}

func eachOptionGroup(group *flags.Group, cb func(*flags.Option)) {
	for _, opt := range group.Options() {
		cb(opt)
	}

	for _, group := range group.Groups() {
		eachOptionGroup(group, cb)
	}
}
