package runners

import "context"

func NetRunner(ctx context.Context) error     { return nil }
func StorRunner(ctx context.Context) error    { return nil }
func VMRunner(ctx context.Context) error      { return nil }
func InfraRRunner(ctx context.Context) error  { return nil }
func PreChkRunner(ctx context.Context) error  { return nil }
func ProvDBRunner(ctx context.Context) error  { return nil }
func SchemaRunner(ctx context.Context) error  { return nil }
func DBRRunner(ctx context.Context) error     { return nil }
func ProvMonRunner(ctx context.Context) error { return nil }
func HookRunner(ctx context.Context) error    { return nil }
func MonRRunner(ctx context.Context) error    { return nil }
func DeployRunner(ctx context.Context) error  { return nil }
func SmokeRunner(ctx context.Context) error   { return nil }
func AppRRunner(ctx context.Context) error    { return nil }
func SecScanRunner(ctx context.Context) error { return nil }
func E2ERunner(ctx context.Context) error     { return nil }
func CommitRunner(ctx context.Context) error  { return nil }
func RollRunner(ctx context.Context) error    { return nil }
func NotifyRunner(ctx context.Context) error  { return nil }
func SinkRunner(ctx context.Context) error    { return nil }
