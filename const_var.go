package main

const (
	helpGrade = "# HELP observatory_http_grade Grade representation of score, A+=12, A=11, A-=10, B+=9, B=8, B-=7, C+=6, C=5, C-=4, D+=3, D=2, D-=1, F=0"
	typeGrade = "# TYPE observatory_http_grade gauge"
	helpScore = "# HELP observatory_http_score Http score"
	typeScore = "# TYPE observatory_http_score gauge"
	helpTest  = "# HELP observatory_http_tests_passed Number of test passed"
	typeTest  = "# TYPE observatory_http_tests_passed gauge"

	helpTLS = "# HELP observatory_tls"
	typeTLS = "# TYPE observatory_tls"

	helptlse     = helpTLS + "_enable Is 1 (aka 'enable') if tls is enable for domain"
	typetlse     = typeTLS + "_enable gauge"
	helptrust    = helpTLS + "_valid Is 1 (aka 'trusted') if certificate is known to be trusted (via truststores)"
	typetrust    = typeTLS + "_valid gauge"
	helplevel    = helpTLS + "_level Defines the Mozilla SSL compatibility level for given domain (0 = old, 1 = bad, 2 = non compliant, 3 = intermediate, 4 = modern)"
	typelevel    = typeTLS + "_level gauge"
	helptlsgrade = helpTLS + "_grade Grade representation of score, A=4, B=3, C=2, D=1, F=0"
	typetlsgrade = typeTLS + "_grade gauge"
	helptlsscore = helpTLS + "_score Defines the score given by Mozilla Observatory's mozillaGradingWorker (0...100)"
	typetlsscore = typeTLS + "_score gauge"

	helpCERT = "# HELP observatory_cert"
	typeCERT = "# TYPE observatory_cert"

	helpexpire = helpCERT + "_expire Expiry date for certificate in timestamp format."
	typeexpire = typeCERT + "_expire gauge"
	helpstart  = helpCERT + "_start Start date for certificate in timestamp format"
	typestart  = typeCERT + "_start gauge"
)
