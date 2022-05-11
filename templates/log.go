package templates

// LogView template for displaying event log in a list.
const LogView = `
{{define "content"}}
	<div class="repository file list">
		<div class="header-wrapper">
			<div class="ui container">
				<div class="ui vertically padded grid head">
					<div class="column">
						<div class="ui header">
							<div class="ui huge breadcrumb">
								<i class="mega-octicon octicon-repo"></i>
							</div>
						</div>
					</div>
				</div>
			</div>
			<div class="ui tabs container">
			</div>
			<div class="ui tabs divider"></div>
		</div>
		<div class="ui container">
			<p id="repo-desc">
			<span class="description has-emoji">Work log</span>
			<a class="link" href=""></a>
			</p>
			<table id="repo-files-table" class="ui unstackable fixed single line table">
				<tbody>
				        {{ $BaseURL := .BaseURL }}
					{{range $job := .Jobs}}
						<tr>
							<td class="name text bold two wide"><a href="{{ $BaseURL }}/log/{{$job.ID}}">Job {{$job.ID}}</a></td>
							<td class="name text bold four wide"><a href="{{ $BaseURL }}/log/{{$job.ID}}">{{$job.Label}}</a></td>
							<td class="name text five wide">{{$job.SubmitTime}}</td>
							<td class="name text five wide">{{$job.EndTime}}</td>
							<td class="name text four wide">{{if $job.Error}}{{$job.Error}}{{end}}</td>
						</tr>
					{{end}}
				</tbody>
			</table>
		</div>
	</div>
{{end}}
`
