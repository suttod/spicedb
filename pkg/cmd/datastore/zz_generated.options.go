// Code generated by github.com/ecordell/optgen. DO NOT EDIT.
package datastore

import "time"

type ConfigOption func(c *Config)

// NewConfigWithOptions creates a new Config with the passed in options set
func NewConfigWithOptions(opts ...ConfigOption) *Config {
	c := &Config{}
	for _, o := range opts {
		o(c)
	}
	return c
}

// ToOption returns a new ConfigOption that sets the values from the passed in Config
func (c *Config) ToOption() ConfigOption {
	return func(to *Config) {
		to.Engine = c.Engine
		to.URI = c.URI
		to.GCWindow = c.GCWindow
		to.LegacyFuzzing = c.LegacyFuzzing
		to.RevisionQuantization = c.RevisionQuantization
		to.MaxRevisionStalenessPercent = c.MaxRevisionStalenessPercent
		to.ReadConnPool = c.ReadConnPool
		to.WriteConnPool = c.WriteConnPool
		to.SplitQueryCount = c.SplitQueryCount
		to.ReadOnly = c.ReadOnly
		to.EnableDatastoreMetrics = c.EnableDatastoreMetrics
		to.DisableStats = c.DisableStats
		to.BootstrapFiles = c.BootstrapFiles
		to.BootstrapFileContents = c.BootstrapFileContents
		to.BootstrapOverwrite = c.BootstrapOverwrite
		to.BootstrapTimeout = c.BootstrapTimeout
		to.RequestHedgingEnabled = c.RequestHedgingEnabled
		to.RequestHedgingInitialSlowValue = c.RequestHedgingInitialSlowValue
		to.RequestHedgingMaxRequests = c.RequestHedgingMaxRequests
		to.RequestHedgingQuantile = c.RequestHedgingQuantile
		to.FollowerReadDelay = c.FollowerReadDelay
		to.MaxRetries = c.MaxRetries
		to.OverlapKey = c.OverlapKey
		to.OverlapStrategy = c.OverlapStrategy
		to.GCInterval = c.GCInterval
		to.GCMaxOperationTime = c.GCMaxOperationTime
		to.SpannerCredentialsFile = c.SpannerCredentialsFile
		to.SpannerEmulatorHost = c.SpannerEmulatorHost
		to.TablePrefix = c.TablePrefix
		to.WatchBufferLength = c.WatchBufferLength
		to.MigrationPhase = c.MigrationPhase
	}
}

// ConfigWithOptions configures an existing Config with the passed in options set
func ConfigWithOptions(c *Config, opts ...ConfigOption) *Config {
	for _, o := range opts {
		o(c)
	}
	return c
}

// WithEngine returns an option that can set Engine on a Config
func WithEngine(engine string) ConfigOption {
	return func(c *Config) {
		c.Engine = engine
	}
}

// WithURI returns an option that can set URI on a Config
func WithURI(uRI string) ConfigOption {
	return func(c *Config) {
		c.URI = uRI
	}
}

// WithGCWindow returns an option that can set GCWindow on a Config
func WithGCWindow(gCWindow time.Duration) ConfigOption {
	return func(c *Config) {
		c.GCWindow = gCWindow
	}
}

// WithLegacyFuzzing returns an option that can set LegacyFuzzing on a Config
func WithLegacyFuzzing(legacyFuzzing time.Duration) ConfigOption {
	return func(c *Config) {
		c.LegacyFuzzing = legacyFuzzing
	}
}

// WithRevisionQuantization returns an option that can set RevisionQuantization on a Config
func WithRevisionQuantization(revisionQuantization time.Duration) ConfigOption {
	return func(c *Config) {
		c.RevisionQuantization = revisionQuantization
	}
}

// WithMaxRevisionStalenessPercent returns an option that can set MaxRevisionStalenessPercent on a Config
func WithMaxRevisionStalenessPercent(maxRevisionStalenessPercent float64) ConfigOption {
	return func(c *Config) {
		c.MaxRevisionStalenessPercent = maxRevisionStalenessPercent
	}
}

// WithReadConnPool returns an option that can set ReadConnPool on a Config
func WithReadConnPool(readConnPool ConnPoolConfig) ConfigOption {
	return func(c *Config) {
		c.ReadConnPool = readConnPool
	}
}

// WithWriteConnPool returns an option that can set WriteConnPool on a Config
func WithWriteConnPool(writeConnPool ConnPoolConfig) ConfigOption {
	return func(c *Config) {
		c.WriteConnPool = writeConnPool
	}
}

// WithSplitQueryCount returns an option that can set SplitQueryCount on a Config
func WithSplitQueryCount(splitQueryCount uint16) ConfigOption {
	return func(c *Config) {
		c.SplitQueryCount = splitQueryCount
	}
}

// WithReadOnly returns an option that can set ReadOnly on a Config
func WithReadOnly(readOnly bool) ConfigOption {
	return func(c *Config) {
		c.ReadOnly = readOnly
	}
}

// WithEnableDatastoreMetrics returns an option that can set EnableDatastoreMetrics on a Config
func WithEnableDatastoreMetrics(enableDatastoreMetrics bool) ConfigOption {
	return func(c *Config) {
		c.EnableDatastoreMetrics = enableDatastoreMetrics
	}
}

// WithDisableStats returns an option that can set DisableStats on a Config
func WithDisableStats(disableStats bool) ConfigOption {
	return func(c *Config) {
		c.DisableStats = disableStats
	}
}

// WithBootstrapFiles returns an option that can append BootstrapFiless to Config.BootstrapFiles
func WithBootstrapFiles(bootstrapFiles string) ConfigOption {
	return func(c *Config) {
		c.BootstrapFiles = append(c.BootstrapFiles, bootstrapFiles)
	}
}

// SetBootstrapFiles returns an option that can set BootstrapFiles on a Config
func SetBootstrapFiles(bootstrapFiles []string) ConfigOption {
	return func(c *Config) {
		c.BootstrapFiles = bootstrapFiles
	}
}

// WithBootstrapFileContents returns an option that can append BootstrapFileContentss to Config.BootstrapFileContents
func WithBootstrapFileContents(key string, value []byte) ConfigOption {
	return func(c *Config) {
		c.BootstrapFileContents[key] = value
	}
}

// SetBootstrapFileContents returns an option that can set BootstrapFileContents on a Config
func SetBootstrapFileContents(bootstrapFileContents map[string][]byte) ConfigOption {
	return func(c *Config) {
		c.BootstrapFileContents = bootstrapFileContents
	}
}

// WithBootstrapOverwrite returns an option that can set BootstrapOverwrite on a Config
func WithBootstrapOverwrite(bootstrapOverwrite bool) ConfigOption {
	return func(c *Config) {
		c.BootstrapOverwrite = bootstrapOverwrite
	}
}

// WithBootstrapTimeout returns an option that can set BootstrapTimeout on a Config
func WithBootstrapTimeout(bootstrapTimeout time.Duration) ConfigOption {
	return func(c *Config) {
		c.BootstrapTimeout = bootstrapTimeout
	}
}

// WithRequestHedgingEnabled returns an option that can set RequestHedgingEnabled on a Config
func WithRequestHedgingEnabled(requestHedgingEnabled bool) ConfigOption {
	return func(c *Config) {
		c.RequestHedgingEnabled = requestHedgingEnabled
	}
}

// WithRequestHedgingInitialSlowValue returns an option that can set RequestHedgingInitialSlowValue on a Config
func WithRequestHedgingInitialSlowValue(requestHedgingInitialSlowValue time.Duration) ConfigOption {
	return func(c *Config) {
		c.RequestHedgingInitialSlowValue = requestHedgingInitialSlowValue
	}
}

// WithRequestHedgingMaxRequests returns an option that can set RequestHedgingMaxRequests on a Config
func WithRequestHedgingMaxRequests(requestHedgingMaxRequests uint64) ConfigOption {
	return func(c *Config) {
		c.RequestHedgingMaxRequests = requestHedgingMaxRequests
	}
}

// WithRequestHedgingQuantile returns an option that can set RequestHedgingQuantile on a Config
func WithRequestHedgingQuantile(requestHedgingQuantile float64) ConfigOption {
	return func(c *Config) {
		c.RequestHedgingQuantile = requestHedgingQuantile
	}
}

// WithFollowerReadDelay returns an option that can set FollowerReadDelay on a Config
func WithFollowerReadDelay(followerReadDelay time.Duration) ConfigOption {
	return func(c *Config) {
		c.FollowerReadDelay = followerReadDelay
	}
}

// WithMaxRetries returns an option that can set MaxRetries on a Config
func WithMaxRetries(maxRetries int) ConfigOption {
	return func(c *Config) {
		c.MaxRetries = maxRetries
	}
}

// WithOverlapKey returns an option that can set OverlapKey on a Config
func WithOverlapKey(overlapKey string) ConfigOption {
	return func(c *Config) {
		c.OverlapKey = overlapKey
	}
}

// WithOverlapStrategy returns an option that can set OverlapStrategy on a Config
func WithOverlapStrategy(overlapStrategy string) ConfigOption {
	return func(c *Config) {
		c.OverlapStrategy = overlapStrategy
	}
}

// WithGCInterval returns an option that can set GCInterval on a Config
func WithGCInterval(gCInterval time.Duration) ConfigOption {
	return func(c *Config) {
		c.GCInterval = gCInterval
	}
}

// WithGCMaxOperationTime returns an option that can set GCMaxOperationTime on a Config
func WithGCMaxOperationTime(gCMaxOperationTime time.Duration) ConfigOption {
	return func(c *Config) {
		c.GCMaxOperationTime = gCMaxOperationTime
	}
}

// WithSpannerCredentialsFile returns an option that can set SpannerCredentialsFile on a Config
func WithSpannerCredentialsFile(spannerCredentialsFile string) ConfigOption {
	return func(c *Config) {
		c.SpannerCredentialsFile = spannerCredentialsFile
	}
}

// WithSpannerEmulatorHost returns an option that can set SpannerEmulatorHost on a Config
func WithSpannerEmulatorHost(spannerEmulatorHost string) ConfigOption {
	return func(c *Config) {
		c.SpannerEmulatorHost = spannerEmulatorHost
	}
}

// WithTablePrefix returns an option that can set TablePrefix on a Config
func WithTablePrefix(tablePrefix string) ConfigOption {
	return func(c *Config) {
		c.TablePrefix = tablePrefix
	}
}

// WithWatchBufferLength returns an option that can set WatchBufferLength on a Config
func WithWatchBufferLength(watchBufferLength uint16) ConfigOption {
	return func(c *Config) {
		c.WatchBufferLength = watchBufferLength
	}
}

// WithMigrationPhase returns an option that can set MigrationPhase on a Config
func WithMigrationPhase(migrationPhase string) ConfigOption {
	return func(c *Config) {
		c.MigrationPhase = migrationPhase
	}
}
