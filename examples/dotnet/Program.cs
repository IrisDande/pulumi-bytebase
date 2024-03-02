using System.Collections.Generic;
using System.Linq;
using Pulumi;
using bytebase = Pulumi.bytebase;

return await Deployment.RunAsync(() => 
{
    var myRandomResource = new bytebase.Random("myRandomResource", new()
    {
        Length = 24,
    });

    return new Dictionary<string, object?>
    {
        ["output"] = 
        {
            { "value", myRandomResource.Result },
        },
    };
});

