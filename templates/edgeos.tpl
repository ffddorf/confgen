interfaces {
  {{- $ifGroups := edgeosPrepareInterfaces .Device.Interfaces }}
  {{- range $ifType, $ifs := $ifGroups }}
  {{- range $ifs }}
  {{ $ifType }} {{ .Name }} {
    {{- if .Description }}
    description {{ .Description | maybeQuote }}
    {{- end }}
    {{- range .Ip_addresses }}
    address {{ .Address }}
    {{- end }}
    {{- if not .Enabled }}
    disable
    {{- end }}
    {{- if .Speed }}
    speed {{ .Speed }}
    {{- end }}
    {{- if .Duplex }}
    duplex {{ .Duplex }}
    {{- end }}
    {{- /* edgeosConfigFromMap .ConfigContext */}}
    {{- range .VIFs }}
    vif {{ .Untagged_vlan.Vid }} {
      {{- if .Description }}
      description {{ .Description | maybeQuote }}
      {{- end }}
      {{- range .Ip_addresses }}
      address {{ .Address }}
      {{- end }}
      {{- if not .Enabled }}
      disable
      {{- end }}
      {{- /* edgeosConfigFromMap .ConfigContext */}}
    }
    {{- end }}
  }
  {{- end }}
  {{- end }}
}
