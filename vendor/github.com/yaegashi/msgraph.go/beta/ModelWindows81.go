// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// Windows81CertificateProfileBase Device Configuration.
type Windows81CertificateProfileBase struct {
	// WindowsCertificateProfileBase is the base model of Windows81CertificateProfileBase
	WindowsCertificateProfileBase
	// ExtendedKeyUsages Extended Key Usage (EKU) settings. This collection can contain a maximum of 500 elements.
	ExtendedKeyUsages []ExtendedKeyUsage `json:"extendedKeyUsages,omitempty"`
	// CustomSubjectAlternativeNames Custom Subject Alternative Name Settings. This collection can contain a maximum of 500 elements.
	CustomSubjectAlternativeNames []CustomSubjectAlternativeName `json:"customSubjectAlternativeNames,omitempty"`
}

// Windows81CompliancePolicy This class contains compliance settings for Windows 8.1.
type Windows81CompliancePolicy struct {
	// DeviceCompliancePolicy is the base model of Windows81CompliancePolicy
	DeviceCompliancePolicy
	// PasswordRequired Require a password to unlock Windows device.
	PasswordRequired *bool `json:"passwordRequired,omitempty"`
	// PasswordBlockSimple Indicates whether or not to block simple password.
	PasswordBlockSimple *bool `json:"passwordBlockSimple,omitempty"`
	// PasswordExpirationDays Password expiration in days.
	PasswordExpirationDays *int `json:"passwordExpirationDays,omitempty"`
	// PasswordMinimumLength The minimum password length.
	PasswordMinimumLength *int `json:"passwordMinimumLength,omitempty"`
	// PasswordMinutesOfInactivityBeforeLock Minutes of inactivity before a password is required.
	PasswordMinutesOfInactivityBeforeLock *int `json:"passwordMinutesOfInactivityBeforeLock,omitempty"`
	// PasswordMinimumCharacterSetCount The number of character sets required in the password.
	PasswordMinimumCharacterSetCount *int `json:"passwordMinimumCharacterSetCount,omitempty"`
	// PasswordRequiredType The required password type.
	PasswordRequiredType *RequiredPasswordType `json:"passwordRequiredType,omitempty"`
	// PasswordPreviousPasswordBlockCount The number of previous passwords to prevent re-use of. Valid values 0 to 24
	PasswordPreviousPasswordBlockCount *int `json:"passwordPreviousPasswordBlockCount,omitempty"`
	// OsMinimumVersion Minimum Windows 8.1 version.
	OsMinimumVersion *string `json:"osMinimumVersion,omitempty"`
	// OsMaximumVersion Maximum Windows 8.1 version.
	OsMaximumVersion *string `json:"osMaximumVersion,omitempty"`
	// StorageRequireEncryption Indicates whether or not to require encryption on a windows 8.1 device.
	StorageRequireEncryption *bool `json:"storageRequireEncryption,omitempty"`
}

// Windows81GeneralConfiguration This topic provides descriptions of the declared methods, properties and relationships exposed by the windows81GeneralConfiguration resource.
type Windows81GeneralConfiguration struct {
	// DeviceConfiguration is the base model of Windows81GeneralConfiguration
	DeviceConfiguration
	// AccountsBlockAddingNonMicrosoftAccountEmail Indicates whether or not to Block the user from adding email accounts to the device that are not associated with a Microsoft account.
	AccountsBlockAddingNonMicrosoftAccountEmail *bool `json:"accountsBlockAddingNonMicrosoftAccountEmail,omitempty"`
	// ApplyOnlyToWindows81 Value indicating whether this policy only applies to Windows 8.1. This property is read-only.
	ApplyOnlyToWindows81 *bool `json:"applyOnlyToWindows81,omitempty"`
	// BrowserBlockAutofill Indicates whether or not to block auto fill.
	BrowserBlockAutofill *bool `json:"browserBlockAutofill,omitempty"`
	// BrowserBlockAutomaticDetectionOfIntranetSites Indicates whether or not to block automatic detection of Intranet sites.
	BrowserBlockAutomaticDetectionOfIntranetSites *bool `json:"browserBlockAutomaticDetectionOfIntranetSites,omitempty"`
	// BrowserBlockEnterpriseModeAccess Indicates whether or not to block enterprise mode access.
	BrowserBlockEnterpriseModeAccess *bool `json:"browserBlockEnterpriseModeAccess,omitempty"`
	// BrowserBlockJavaScript Indicates whether or not to Block the user from using JavaScript.
	BrowserBlockJavaScript *bool `json:"browserBlockJavaScript,omitempty"`
	// BrowserBlockPlugins Indicates whether or not to block plug-ins.
	BrowserBlockPlugins *bool `json:"browserBlockPlugins,omitempty"`
	// BrowserBlockPopups Indicates whether or not to block popups.
	BrowserBlockPopups *bool `json:"browserBlockPopups,omitempty"`
	// BrowserBlockSendingDoNotTrackHeader Indicates whether or not to Block the user from sending the do not track header.
	BrowserBlockSendingDoNotTrackHeader *bool `json:"browserBlockSendingDoNotTrackHeader,omitempty"`
	// BrowserBlockSingleWordEntryOnIntranetSites Indicates whether or not to block a single word entry on Intranet sites.
	BrowserBlockSingleWordEntryOnIntranetSites *bool `json:"browserBlockSingleWordEntryOnIntranetSites,omitempty"`
	// BrowserRequireSmartScreen Indicates whether or not to require the user to use the smart screen filter.
	BrowserRequireSmartScreen *bool `json:"browserRequireSmartScreen,omitempty"`
	// BrowserEnterpriseModeSiteListLocation The enterprise mode site list location. Could be a local file, local network or http location.
	BrowserEnterpriseModeSiteListLocation *string `json:"browserEnterpriseModeSiteListLocation,omitempty"`
	// BrowserInternetSecurityLevel The internet security level.
	BrowserInternetSecurityLevel *InternetSiteSecurityLevel `json:"browserInternetSecurityLevel,omitempty"`
	// BrowserIntranetSecurityLevel The Intranet security level.
	BrowserIntranetSecurityLevel *SiteSecurityLevel `json:"browserIntranetSecurityLevel,omitempty"`
	// BrowserLoggingReportLocation The logging report location.
	BrowserLoggingReportLocation *string `json:"browserLoggingReportLocation,omitempty"`
	// BrowserRequireHighSecurityForRestrictedSites Indicates whether or not to require high security for restricted sites.
	BrowserRequireHighSecurityForRestrictedSites *bool `json:"browserRequireHighSecurityForRestrictedSites,omitempty"`
	// BrowserRequireFirewall Indicates whether or not to require a firewall.
	BrowserRequireFirewall *bool `json:"browserRequireFirewall,omitempty"`
	// BrowserRequireFraudWarning Indicates whether or not to require fraud warning.
	BrowserRequireFraudWarning *bool `json:"browserRequireFraudWarning,omitempty"`
	// BrowserTrustedSitesSecurityLevel The trusted sites security level.
	BrowserTrustedSitesSecurityLevel *SiteSecurityLevel `json:"browserTrustedSitesSecurityLevel,omitempty"`
	// CellularBlockDataRoaming Indicates whether or not to block data roaming.
	CellularBlockDataRoaming *bool `json:"cellularBlockDataRoaming,omitempty"`
	// DiagnosticsBlockDataSubmission Indicates whether or not to block diagnostic data submission.
	DiagnosticsBlockDataSubmission *bool `json:"diagnosticsBlockDataSubmission,omitempty"`
	// PasswordBlockPicturePasswordAndPin Indicates whether or not to Block the user from using a pictures password and pin.
	PasswordBlockPicturePasswordAndPin *bool `json:"passwordBlockPicturePasswordAndPin,omitempty"`
	// PasswordExpirationDays Password expiration in days.
	PasswordExpirationDays *int `json:"passwordExpirationDays,omitempty"`
	// PasswordMinimumLength The minimum password length.
	PasswordMinimumLength *int `json:"passwordMinimumLength,omitempty"`
	// PasswordMinutesOfInactivityBeforeScreenTimeout The minutes of inactivity before the screen times out.
	PasswordMinutesOfInactivityBeforeScreenTimeout *int `json:"passwordMinutesOfInactivityBeforeScreenTimeout,omitempty"`
	// PasswordMinimumCharacterSetCount The number of character sets required in the password.
	PasswordMinimumCharacterSetCount *int `json:"passwordMinimumCharacterSetCount,omitempty"`
	// PasswordPreviousPasswordBlockCount The number of previous passwords to prevent re-use of. Valid values 0 to 24
	PasswordPreviousPasswordBlockCount *int `json:"passwordPreviousPasswordBlockCount,omitempty"`
	// PasswordRequiredType The required password type.
	PasswordRequiredType *RequiredPasswordType `json:"passwordRequiredType,omitempty"`
	// PasswordSignInFailureCountBeforeFactoryReset The number of sign in failures before factory reset.
	PasswordSignInFailureCountBeforeFactoryReset *int `json:"passwordSignInFailureCountBeforeFactoryReset,omitempty"`
	// StorageRequireDeviceEncryption Indicates whether or not to require encryption on a mobile device.
	StorageRequireDeviceEncryption *bool `json:"storageRequireDeviceEncryption,omitempty"`
	// MinimumAutoInstallClassification The minimum update classification to install automatically.
	MinimumAutoInstallClassification *UpdateClassification `json:"minimumAutoInstallClassification,omitempty"`
	// UpdatesMinimumAutoInstallClassification The minimum update classification to install automatically.
	UpdatesMinimumAutoInstallClassification *UpdateClassification `json:"updatesMinimumAutoInstallClassification,omitempty"`
	// UpdatesRequireAutomaticUpdates Indicates whether or not to require automatic updates.
	UpdatesRequireAutomaticUpdates *bool `json:"updatesRequireAutomaticUpdates,omitempty"`
	// UserAccountControlSettings The user account control settings.
	UserAccountControlSettings *WindowsUserAccountControlSettings `json:"userAccountControlSettings,omitempty"`
	// WorkFoldersURL The work folders url.
	WorkFoldersURL *string `json:"workFoldersUrl,omitempty"`
}

// Windows81SCEPCertificateProfile Windows 8.1+ SCEP certificate profile
type Windows81SCEPCertificateProfile struct {
	// Windows81CertificateProfileBase is the base model of Windows81SCEPCertificateProfile
	Windows81CertificateProfileBase
	// ScepServerUrls SCEP Server Url(s).
	ScepServerUrls []string `json:"scepServerUrls,omitempty"`
	// SubjectNameFormatString Custom format to use with SubjectNameFormat = Custom. Example: CN={{EmailAddress}},E={{EmailAddress}},OU=Enterprise Users,O=Contoso Corporation,L=Redmond,ST=WA,C=US
	SubjectNameFormatString *string `json:"subjectNameFormatString,omitempty"`
	// KeyUsage SCEP Key Usage.
	KeyUsage *KeyUsages `json:"keyUsage,omitempty"`
	// KeySize SCEP Key Size.
	KeySize *KeySize `json:"keySize,omitempty"`
	// HashAlgorithm SCEP Hash Algorithm.
	HashAlgorithm *HashAlgorithms `json:"hashAlgorithm,omitempty"`
	// SubjectAlternativeNameFormatString Custom String that defines the AAD Attribute.
	SubjectAlternativeNameFormatString *string `json:"subjectAlternativeNameFormatString,omitempty"`
	// CertificateStore Target store certificate
	CertificateStore *CertificateStore `json:"certificateStore,omitempty"`
	// RootCertificate undocumented
	RootCertificate *Windows81TrustedRootCertificate `json:"rootCertificate,omitempty"`
	// ManagedDeviceCertificateStates undocumented
	ManagedDeviceCertificateStates []ManagedDeviceCertificateState `json:"managedDeviceCertificateStates,omitempty"`
}

// Windows81TrustedRootCertificate Windows 8.1 Trusted Certificate configuration profile
type Windows81TrustedRootCertificate struct {
	// DeviceConfiguration is the base model of Windows81TrustedRootCertificate
	DeviceConfiguration
	// TrustedRootCertificate Trusted Root Certificate
	TrustedRootCertificate *Binary `json:"trustedRootCertificate,omitempty"`
	// CertFileName File name to display in UI.
	CertFileName *string `json:"certFileName,omitempty"`
	// DestinationStore Destination store location for the Trusted Root Certificate.
	DestinationStore *CertificateDestinationStore `json:"destinationStore,omitempty"`
}

// Windows81VpnConfiguration By providing the configurations in this profile you can instruct the Windows 8.1 (and later) devices to connect to desired VPN endpoint. By specifying the authentication method and security types expected by VPN endpoint you can make the VPN connection seamless for end user.
type Windows81VpnConfiguration struct {
	// WindowsVPNConfiguration is the base model of Windows81VpnConfiguration
	WindowsVPNConfiguration
	// ApplyOnlyToWindows81 Value indicating whether this policy only applies to Windows 8.1. This property is read-only.
	ApplyOnlyToWindows81 *bool `json:"applyOnlyToWindows81,omitempty"`
	// ConnectionType Connection type.
	ConnectionType *WindowsVPNConnectionType `json:"connectionType,omitempty"`
	// LoginGroupOrDomain Login group or domain when connection type is set to Dell SonicWALL Mobile Connection.
	LoginGroupOrDomain *string `json:"loginGroupOrDomain,omitempty"`
	// EnableSplitTunneling Enable split tunneling for the VPN.
	EnableSplitTunneling *bool `json:"enableSplitTunneling,omitempty"`
	// ProxyServer Proxy Server.
	ProxyServer *Windows81VpnProxyServer `json:"proxyServer,omitempty"`
}

// Windows81VpnProxyServer undocumented
type Windows81VpnProxyServer struct {
	// VPNProxyServer is the base model of Windows81VpnProxyServer
	VPNProxyServer
	// AutomaticallyDetectProxySettings Automatically detect proxy settings.
	AutomaticallyDetectProxySettings *bool `json:"automaticallyDetectProxySettings,omitempty"`
	// BypassProxyServerForLocalAddress Bypass proxy server for local address.
	BypassProxyServerForLocalAddress *bool `json:"bypassProxyServerForLocalAddress,omitempty"`
}

// Windows81WifiImportConfiguration Windows 8.1+ Wi-Fi import configuration. By configuring this profile you can instruct Windows 8.1 (and later) devices to connect to desired Wi-Fi endpoint. Connect a Windows 8.1 device to the desired Wi-Fi network and extract the XML from that device to later embed into this Wi-Fi profile.
type Windows81WifiImportConfiguration struct {
	// DeviceConfiguration is the base model of Windows81WifiImportConfiguration
	DeviceConfiguration
	// PayloadFileName Payload file name (*.xml).
	PayloadFileName *string `json:"payloadFileName,omitempty"`
	// ProfileName Profile name displayed in the UI.
	ProfileName *string `json:"profileName,omitempty"`
	// Payload Payload. (UTF8 encoded byte array). This is the XML file saved on the device you used to connect to the Wi-Fi endpoint.
	Payload *Binary `json:"payload,omitempty"`
}
