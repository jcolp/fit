// Generated by reader_util_test.go.
// Only edit to bootstrap new entries or change existing entries.

package fit_test

var decodeTestFiles = [...]struct {
	folder      string
	name        string
	source      string
	wantErr     bool
	fingerprint uint64
	compress    bool
	dopts       testingDecodeOpts
	skipEncode  bool
	encodeNote  string
}{
	{
		"me",
		"activity-small-fenix2-run.fit",
		"",
		false,
		18371229752680076200,
		true,
		tdoAllWithDiscardLogger,
		true,
		"Decode mismatch due to unknown fields",
	},
	{
		"fitsdk",
		"Activity.fit",
		"",
		false,
		14144875446989160045,
		true,
		tdoNone,
		false,
		"",
	},
	{
		"fitsdk",
		"MonitoringFile.fit",
		"",
		false,
		11936585269915402423,
		true,
		tdoNone,
		true,
		"Fails because first message has different valid fields (#36)",
	},
	{
		"fitsdk",
		"Settings.fit",
		"",
		false,
		18422634047426156243,
		true,
		tdoNone,
		false,
		"",
	},
	{
		"fitsdk",
		"WeightScaleMultiUser.fit",
		"",
		false,
		3311688825924049582,
		true,
		tdoNone,
		false,
		"",
	},
	{
		"fitsdk",
		"WorkoutCustomTargetValues.fit",
		"",
		false,
		11173674486458549276,
		true,
		tdoNone,
		false,
		"",
	},
	{
		"fitsdk",
		"WorkoutIndividualSteps.fit",
		"",
		false,
		13551116123795560734,
		true,
		tdoNone,
		false,
		"",
	},
	{
		"fitsdk",
		"WorkoutRepeatGreaterThanStep.fit",
		"",
		false,
		551789846658266529,
		true,
		tdoNone,
		false,
		"",
	},
	{
		"fitsdk",
		"WorkoutRepeatSteps.fit",
		"",
		false,
		13365155552315550997,
		true,
		tdoNone,
		false,
		"",
	},
	{
		"fitsdk",
		"WeightScaleSingleUser.fit",
		"",
		false,
		4152838174762136445,
		true,
		tdoNone,
		false,
		"",
	},
	{
		"python-fitparse",
		"garmin-edge-500-activitiy.fit",
		"",
		false,
		12438793213158767420,
		true,
		tdoNone,
		false,
		"",
	},
	{
		"python-fitparse",
		"sample-activity-indoor-trainer.fit",
		"",
		false,
		14338690625832646283,
		true,
		tdoNone,
		true,
		"Fails because first message has different valid fields (#36)",
	},
	{
		"python-fitparse",
		"compressed-speed-distance.fit",
		"",
		false,
		0,
		false,
		tdoNone,
		false,
		"",
	},
	{
		"python-fitparse",
		"antfs-dump.63.fit",
		"",
		false,
		6582447834165183736,
		true,
		tdoNone,
		true,
		"Fails because first message has different valid fields (#36)",
	},
	{
		"sram",
		"Settings.fit",
		"",
		false,
		5866657363356029809,
		true,
		tdoNone,
		true,
		"Fails due to encoder using profile length for arrays (#37)",
	},
	{
		"sram",
		"Settings2.fit",
		"",
		false,
		15709312684722569429,
		true,
		tdoNone,
		true,
		"Fails due to encoder using profile length for arrays (#37)",
	},
	{
		"dcrainmaker",
		"Edge810-Vector-2013-08-16-15-35-10.fit",
		"",
		false,
		17721372148334979063,
		true,
		tdoNone,
		true,
		"Fails because first message has different valid fields (#36)",
	},
	{
		"misc",
		"2013-02-06-12-11-14.fit",
		"",
		false,
		5737379589684728556,
		true,
		tdoNone,
		false,
		"",
	},
	{
		"misc",
		"2015-10-13-08-43-15.fit",
		"",
		false,
		1804560237650400088,
		true,
		tdoNone,
		false,
		"",
	},
	{
		"bpg",
		"garmin.fit",
		"https://github.com/tormoder/fit/pull/54",
		false,
		9530971717682041541,
		true,
		tdoNone,
		false,
		"Has a definition message with number of fields >85",
	},
	{
		"corrupt",
		"activity-filecrc.fit",
		"",
		true,
		2764261087675073317,
		true,
		tdoNone,
		false,
		"",
	},
	{
		"corrupt",
		"activity-unexpected-eof.fit",
		"",
		true,
		15370168455200986360,
		true,
		tdoNone,
		false,
		"",
	},
	{
		"misc",
		"0134902991.fit",
		"https://github.com/tormoder/fit/pull/59",
		false,
		5389817594289548806,
		true,
		tdoNone,
		true,
		"Contains developer data fields",
	},
	{
		"misc",
		"mornindew-broken.fit",
		"https://github.com/tormoder/fit/issues/41",
		false,
		1206787239575569203,
		true,
		tdoNone,
		true,
		"Contains developer data fields",
	},
	{
		"fitsdk",
		"DeveloperData.fit",
		"",
		false,
		15008309215243872755,
		true,
		tdoNone,
		true,
		"Contains developer data fields",
	},
}
