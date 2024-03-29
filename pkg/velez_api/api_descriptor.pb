
�
grpc/velez_api.proto	velez_apigoogle/api/annotations.protovalidate/validate.protogoogle/protobuf/timestamp.proto":
Version	
Request$
Response
version (	Rversion"�
PortBindings
host (Rhost
	container (R	container8
protoc (2 .velez_api.PortBindings.ProtocolRprotoc")
Protocol
unknown 
tcp
udp"B
VolumeBindings
host (	Rhost
	container (	R	container"/
Image
name (	Rname
tags (	Rtags"�
Smerd
uuid (	Ruuid
name (	Rname

image_name (	R	imageName-
ports (2.velez_api.PortBindingsRports3
volumes (2.velez_api.VolumeBindingsRvolumes/
status (2.velez_api.Smerd.StatusRstatus9

created_at (2.google.protobuf.TimestampR	createdAt"o
Status
unknown 
created

restarting
running
removing

paused

exited
dead"�
	Container�
Hardware"

cpu_amount (H R	cpuAmount�
ram_mb (HRramMb�)
memory_swap_mb (HRmemorySwapMb�B
_cpu_amountB	
_ram_mbB
_memory_swap_mbn
Settings-
ports (2.velez_api.PortBindingsRports3
volumes (2.velez_api.VolumeBindingsRvolumes"�
CreateSmerd�
Request
name (	RnameG

image_name (	B(�B%r#(22([a-z]+)/([a-z-]+):([a-z0-9.]+)R	imageName>
hardware (2.velez_api.Container.HardwareH Rhardware�>
settings (2.velez_api.Container.SettingsHRsettings�)
allow_duplicates (RallowDuplicates
command (	HRcommand�B
	_hardwareB
	_settingsB

_command"�

ListSmerds�
Request
limit (H Rlimit�
name (	HRname�*
general_search (	HRgeneralSearch�
id (	HRid�B
_limitB
_nameB
_general_searchB
_id4
Response(
smerds (2.velez_api.SmerdRsmerds"�
	DropSmerd3
Request
uuids (	Ruuids
name (	Rname�
Response;
failed (2#.velez_api.DropSmerd.Response.ErrorRfailed

successful (	R
successful1
Error
uuid (	Ruuid
cause (	Rcause"�
GetHardware	
Request�
Response7
cpu (2%.velez_api.GetHardware.Response.ValueRcpu@
disk_mem (2%.velez_api.GetHardware.Response.ValueRdiskMem7
ram (2%.velez_api.GetHardware.Response.ValueRram'
ports_available (RportsAvailable%
ports_occupied (RportsOccupied/
Value
value (	Rvalue
err (	Rerr2�
VelezAPIW
Version.velez_api.Version.Request.velez_api.Version.Response"���"/version:*Y
CreateSmerd.velez_api.CreateSmerd.Request.velez_api.Smerd"���"/smerd/create:*c

ListSmerds.velez_api.ListSmerds.Request.velez_api.ListSmerds.Response"���"/smerd/list:*`
	DropSmerd.velez_api.DropSmerd.Request.velez_api.DropSmerd.Response"���"/smerd/drop:*a
GetHardware.velez_api.GetHardware.Request.velez_api.GetHardware.Response"���	/hardwareBZ
/velez_apibproto3