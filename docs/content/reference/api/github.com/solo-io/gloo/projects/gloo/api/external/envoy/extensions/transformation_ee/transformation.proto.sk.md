
---
title: "transformation.proto"
weight: 5
---

<!-- Code generated by solo-kit. DO NOT EDIT. -->


### Package: `envoy.config.filter.http.transformation_ee.v2` 
#### Types:


- [FilterTransformations](#filtertransformations)
- [TransformationRule](#transformationrule)
- [RouteTransformations](#routetransformations)
- [Transformation](#transformation)
- [DlpTransformation](#dlptransformation)
- [Action](#action)
  



##### Source File: [github.com/solo-io/gloo/projects/gloo/api/external/envoy/extensions/transformation_ee/transformation.proto](https://github.com/solo-io/gloo/blob/master/projects/gloo/api/external/envoy/extensions/transformation_ee/transformation.proto)





---
### FilterTransformations



```yaml
"transformations": []envoy.config.filter.http.transformation_ee.v2.TransformationRule

```

| Field | Type | Description | Default |
| ----- | ---- | ----------- |----------- | 
| `transformations` | [[]envoy.config.filter.http.transformation_ee.v2.TransformationRule](../transformation.proto.sk/#transformationrule) | Specifies transformations based on the route matches. The first matched transformation will be applied. If there are overlapped match conditions, please put the most specific match first. |  |




---
### TransformationRule



```yaml
"match": .envoy.api.v2.route.RouteMatch
"routeTransformations": .envoy.config.filter.http.transformation_ee.v2.RouteTransformations

```

| Field | Type | Description | Default |
| ----- | ---- | ----------- |----------- | 
| `match` | [.envoy.api.v2.route.RouteMatch](../../../../../../../../../../../envoy/api/v2/route/route.proto.sk/#routematch) | The route matching parameter. Only when the match is satisfied, the "requires" field will apply. For example: following match will match all requests. .. code-block:: yaml match: prefix: /. |  |
| `routeTransformations` | [.envoy.config.filter.http.transformation_ee.v2.RouteTransformations](../transformation.proto.sk/#routetransformations) | transformation to perform. |  |




---
### RouteTransformations



```yaml
"requestTransformation": .envoy.config.filter.http.transformation_ee.v2.Transformation
"clearRouteCache": bool
"responseTransformation": .envoy.config.filter.http.transformation_ee.v2.Transformation

```

| Field | Type | Description | Default |
| ----- | ---- | ----------- |----------- | 
| `requestTransformation` | [.envoy.config.filter.http.transformation_ee.v2.Transformation](../transformation.proto.sk/#transformation) |  |  |
| `clearRouteCache` | `bool` | clear the route cache if the request transformation was applied. |  |
| `responseTransformation` | [.envoy.config.filter.http.transformation_ee.v2.Transformation](../transformation.proto.sk/#transformation) |  |  |




---
### Transformation



```yaml
"dlpTransformation": .envoy.config.filter.http.transformation_ee.v2.DlpTransformation

```

| Field | Type | Description | Default |
| ----- | ---- | ----------- |----------- | 
| `dlpTransformation` | [.envoy.config.filter.http.transformation_ee.v2.DlpTransformation](../transformation.proto.sk/#dlptransformation) |  |  |




---
### DlpTransformation



```yaml
"actions": []envoy.config.filter.http.transformation_ee.v2.Action

```

| Field | Type | Description | Default |
| ----- | ---- | ----------- |----------- | 
| `actions` | [[]envoy.config.filter.http.transformation_ee.v2.Action](../transformation.proto.sk/#action) | list of actions to apply. |  |




---
### Action



```yaml
"name": string
"regex": []string
"shadow": bool
"percent": .envoy.type.Percent
"maskChar": string

```

| Field | Type | Description | Default |
| ----- | ---- | ----------- |----------- | 
| `name` | `string` | Identifier for this action. Used mostly to help ID specific actions in logs. If left null will default to unknown. |  |
| `regex` | `[]string` | List of regexes to apply to the response body to match data which should be masked They will be applied iteratively in the order which they are specified. |  |
| `shadow` | `bool` | If specified, this rule will not actually be applied, but only logged. |  |
| `percent` | [.envoy.type.Percent](../../../../../../../../../../../envoy/type/percent.proto.sk/#percent) | The percent of the string which should be masked. If not set, defaults to 75%. |  |
| `maskChar` | `string` | The character which should overwrite the masked data If left empty, defaults to "X". |  |





<!-- Start of HubSpot Embed Code -->
<script type="text/javascript" id="hs-script-loader" async defer src="//js.hs-scripts.com/5130874.js"></script>
<!-- End of HubSpot Embed Code -->
